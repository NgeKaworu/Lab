package com.android.example.camera2

import android.Manifest
import android.content.Context
import android.content.pm.PackageManager
import android.graphics.ImageFormat
import android.hardware.camera2.*
import android.hardware.camera2.params.OutputConfiguration
import android.hardware.camera2.params.SessionConfiguration
import android.media.Image
import android.media.ImageReader
import android.os.Build
import android.os.Handler
import android.os.HandlerThread
import android.util.Log
import android.view.Surface
import androidx.annotation.RequiresApi
import androidx.core.app.ActivityCompat
import java.nio.ByteBuffer
import java.nio.ReadOnlyBufferException
import kotlin.experimental.inv

enum class WorkingOutput {
    FRAME_PROCESS,
    RECORDING,
    PREVIEW
}

class Camera2Util(context: Context) {
    private var ctx: Context = context

    // 获取相机manager
    private val cameraManager = context.getSystemService(Context.CAMERA_SERVICE) as CameraManager

    //镜头id
    private val cameraId = cameraManager.cameraIdList[0]

    private var cameraDevice: CameraDevice? = null

    private var cameraCaptureSession: CameraCaptureSession? = null

    private var worker = mutableSetOf<WorkingOutput>()

    // 画面处理器
    private var imageReader = ImageReader.newInstance(640, 480, ImageFormat.YUV_420_888, 2)

    private var previewSurface: Surface? = null;

    /** [HandlerThread] where all camera operations run */
    private val cameraThread = HandlerThread("CameraThread").apply { start() }
    private val cameraHandler = Handler(cameraThread.looper)

    private fun openCamera(callback: (camera: CameraDevice?) -> Unit) {
        // 开启摄像头
        if (ActivityCompat.checkSelfPermission(ctx,
                Manifest.permission.CAMERA) != PackageManager.PERMISSION_GRANTED
        ) {
            // TODO: Consider calling
            //    ActivityCompat#requestPermissions
            // here to request the missing permissions, and then overriding
            //   public void onRequestPermissionsResult(int requestCode, String[] permissions,
            //                                          int[] grantResults)
            // to handle the case where the user grants the permission. See the documentation
            // for ActivityCompat#requestPermissions for more details.
            return
        }

        val characteristics = cameraManager.getCameraCharacteristics(cameraId);
        val level = characteristics.get(CameraCharacteristics.INFO_SUPPORTED_HARDWARE_LEVEL)
//        val outputSize = characteristics.get(CameraCharacteristics.SCALER_STREAM_CONFIGURATION_MAP)
//        val l = outputSize?.getOutputSizes(SurfaceTexture::class.java)
        Log.i(TAG, "openCamera: $level")

        cameraManager.openCamera(cameraId, object : CameraDevice.StateCallback() {
            override fun onOpened(camera: CameraDevice) {
                Log.d(TAG, "camera onOpened")
                cameraDevice = camera
                callback(camera)
            }

            override fun onDisconnected(camera: CameraDevice) {
                Log.d(TAG, "camera onDisconnected")
                camera.close()
                cameraDevice = null
                callback(null)
            }

            override fun onError(camera: CameraDevice, error: Int) {
                Log.d(TAG, "camera onError")
                camera.close()
                cameraDevice = null
                callback(null)
            }
        }, cameraHandler)
    }

    fun closeCamera() {
        if (cameraDevice != null) {
            cameraDevice!!.close()
            cameraDevice = null
        }
    }


    @RequiresApi(Build.VERSION_CODES.P)
    private fun createCaptureSession(
        camera: CameraDevice,
        callback: (session: CameraCaptureSession?) -> Unit,
    ) {
        camera.createCaptureSession(SessionConfiguration(SessionConfiguration.SESSION_REGULAR,
            processOutput(),
            { it.run() },
            object : CameraCaptureSession.StateCallback() {
                override fun onConfigured(session: CameraCaptureSession) {
                    Log.i(TAG, "onConfigured: ")
                    cameraCaptureSession = session
                    callback(session)
                }

                override fun onConfigureFailed(session: CameraCaptureSession) {
                    Log.e(TAG, "onConfigureFailed: ")
                    session.close()
                    callback(null)
                    cameraCaptureSession = null
                }
            })
        )
    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun prepare(callback: (session: CameraCaptureSession?) -> Unit) {
        openCamera { camera -> camera?.apply { createCaptureSession(camera, callback) } }
    }

    fun isPrepared(): Boolean {
        return cameraCaptureSession != null && cameraDevice != null
    }


    @RequiresApi(Build.VERSION_CODES.P)
    private fun reConfiguration(callback: (session: CameraCaptureSession?) -> Unit) {
        cameraCaptureSession?.close()
        cameraDevice?.let { createCaptureSession(it, callback) }
    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun startPreview(surface: Surface) {
        previewSurface = surface
        if (!inPreviewing()) {
            worker.add(WorkingOutput.PREVIEW)
            reConfiguration { startCapture() }
        } else {
            startCapture()
        }


    }

    fun stopPreview() {
        worker.remove(WorkingOutput.PREVIEW)
        continueCapture()
    }

    fun inPreviewing(): Boolean {
        return worker.contains(WorkingOutput.PREVIEW)
    }

    fun setImageReader(ir: ImageReader) {
        imageReader = ir
    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun startScan(listener: ImageReader.OnImageAvailableListener) {
        imageReader.setOnImageAvailableListener(listener, cameraHandler)
        if (!inScanning()) {
            worker.add(WorkingOutput.FRAME_PROCESS)
            reConfiguration { startCapture() }
        } else {
            startCapture()
        }
    }

    fun stopScan() {
        worker.remove(WorkingOutput.FRAME_PROCESS)
        continueCapture()
    }


    fun inScanning(): Boolean {
        return worker.contains(WorkingOutput.FRAME_PROCESS)
    }

    fun release() {
        cameraThread.quitSafely()
    }

    private fun startCapture() {
        // 开始捕获
        cameraCaptureSession!!.setRepeatingBurst(processRequests(), null, cameraHandler)
    }

    private fun continueCapture() {
        if (worker.isEmpty()) {
            cameraCaptureSession?.stopRepeating()
        } else {
            startCapture()
        }
    }

    private fun processRequests(): List<CaptureRequest> {
        val list = mutableListOf<CaptureRequest>()
        if (worker.contains(WorkingOutput.FRAME_PROCESS)) {
            val frameRequest =
                cameraCaptureSession!!.device.createCaptureRequest(CameraDevice.TEMPLATE_STILL_CAPTURE)
                    .apply { addTarget(imageReader.surface) }
                    .build()

            list.add(frameRequest)

        }
        if (worker.contains(WorkingOutput.PREVIEW)) {
            val previewRequest =
                cameraCaptureSession!!.device.createCaptureRequest(CameraDevice.TEMPLATE_PREVIEW)
                    .apply { addTarget(previewSurface!!) }
                    .build()

            list.add(previewRequest)
        }
//        if (worker.contains(WorkingOutput.RECORDING)) {
//        }

        return list
    }

    @RequiresApi(Build.VERSION_CODES.N)
    private fun processOutput(): List<OutputConfiguration> {
        val list = mutableListOf<OutputConfiguration>()

        if (worker.contains(WorkingOutput.FRAME_PROCESS)) {
            list.add(OutputConfiguration(imageReader.surface))
        }

        if (worker.contains(WorkingOutput.PREVIEW)) {
            list.add(OutputConfiguration(previewSurface!!))
        }
//        if (worker.contains(WorkingOutput.RECORDING)) {
//        }

        return list
    }


    companion object {
        private const val TAG = "FaceScannerTAG"

        fun YUV_420_888toNV21(image: Image): ByteArray {
            val width: Int = image.width
            val height: Int = image.height
            val ySize = width * height
            val uvSize = width * height / 4
            val nv21 = ByteArray(ySize + uvSize * 2)
            val yBuffer: ByteBuffer = image.planes[0].buffer // Y
            val uBuffer: ByteBuffer = image.planes[1].buffer // U
            val vBuffer: ByteBuffer = image.planes[2].buffer // V
            var rowStride: Int = image.planes[0].rowStride
            assert(image.planes[0].pixelStride == 1)
            var pos = 0
            if (rowStride == width) { // likely
                yBuffer.get(nv21, 0, ySize)
                pos += ySize
            } else {
                var yBufferPos = -rowStride // not an actual position
                while (pos < ySize) {
                    yBufferPos += rowStride
                    yBuffer.position(yBufferPos)
                    yBuffer.get(nv21, pos, width)
                    pos += width
                }
            }
            rowStride = image.planes.get(2).rowStride
            val pixelStride: Int = image.planes.get(2).pixelStride
            assert(rowStride == image.planes.get(1).rowStride)
            assert(pixelStride == image.planes.get(1).pixelStride)
            if (pixelStride == 2 && rowStride == width && uBuffer.get(0) == vBuffer.get(1)) {
                // maybe V an U planes overlap as per NV21, which means vBuffer[1] is alias of uBuffer[0]
                val savePixel: Byte = vBuffer.get(1)
                try {
                    vBuffer.put(1, savePixel.inv())
                    if (uBuffer.get(0) == savePixel.inv()) {
                        vBuffer.put(1, savePixel)
                        vBuffer.position(0)
                        uBuffer.position(0)
                        vBuffer.get(nv21, ySize, 1)
                        uBuffer.get(nv21, ySize + 1, uBuffer.remaining())
                        return nv21 // shortcut
                    }
                } catch (ex: ReadOnlyBufferException) {
                    // unfortunately, we cannot check if vBuffer and uBuffer overlap
                }

                // unfortunately, the check failed. We must save U and V pixel by pixel
                vBuffer.put(1, savePixel)
            }

            // other optimizations could check if (pixelStride == 1) or (pixelStride == 2),
            // but performance gain would be less significant
            for (row in 0 until height / 2) {
                for (col in 0 until width / 2) {
                    val vuPos = col * pixelStride + row * rowStride
                    nv21[pos++] = vBuffer.get(vuPos)
                    nv21[pos++] = uBuffer.get(vuPos)
                }
            }
            return nv21
        }
    }
}
