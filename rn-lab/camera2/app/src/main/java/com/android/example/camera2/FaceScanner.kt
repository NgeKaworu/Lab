package com.android.example.camera2

import android.Manifest
import android.content.Context
import android.content.pm.PackageManager
import android.graphics.ImageFormat
import android.hardware.camera2.CameraCaptureSession
import android.hardware.camera2.CameraDevice
import android.hardware.camera2.CameraManager
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
import kotlin.coroutines.Continuation
import kotlin.experimental.inv

class FaceScanner(context: Context) {
    private var ctx: Context = context

    // 获取相机manager
    private val cameraManager = context.getSystemService(Context.CAMERA_SERVICE) as CameraManager

    //镜头id
    private val cameraId = cameraManager.cameraIdList[0]

    private var cameraDevice: CameraDevice? = null

    private var _session: CameraCaptureSession? = null


    // 镜头消费者
    private val imageReader = ImageReader.newInstance(640, 480, ImageFormat.YUV_420_888, 1)


    private val targets: MutableList<Surface> = mutableListOf(
//        imageReader.surface
    )


    /** [HandlerThread] where all camera operations run */
    private val cameraThread = HandlerThread("CameraThread").apply { start() }
    private val cameraHandler = Handler(cameraThread.looper)

    fun toggleCamera(callback: (camera: CameraDevice?) -> Unit) {
        return if (cameraDevice == null) {
            openCamera(callback)
        } else {
            closeCamera()
        }
    }

    fun openCamera(callback: (camera: CameraDevice?) -> Unit) {
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

    private fun closeCamera() {
        if (cameraDevice != null) {
            cameraDevice!!.close()
            cameraDevice = null
        }

        if (targets.size > 0) {
            targets.clear()
        }
    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun createCaptureSession(
        camera: CameraDevice,
        callback: (session: CameraCaptureSession?) -> Unit,
    ) {
        val cfg = SessionConfiguration(SessionConfiguration.SESSION_REGULAR,
            targets.map { surface ->
                OutputConfiguration(surface)
            },
            { it.run() },
            object : CameraCaptureSession.StateCallback() {
                override fun onConfigured(session: CameraCaptureSession) {
                    Log.i(TAG, "onConfigured: ")
                    _session = session
                    callback(session)
                }

                override fun onConfigureFailed(session: CameraCaptureSession) {
                    Log.e(TAG, "onConfigureFailed: ")
                    session.close()
                    callback(null)
                    _session = null
                }
            })


        camera.createCaptureSession(cfg)
    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun addSurface(surface: Surface) {
        targets.add(surface)
    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun startScan(width: Int, height: Int, listener: ImageReader.OnImageAvailableListener) {
        // 镜头消费者
        val imageReader = ImageReader.newInstance(width, height, ImageFormat.YUV_420_888, 1).apply {
            setOnImageAvailableListener(listener, cameraHandler)
        }

        // 输出配置
        val cfg = OutputConfiguration(imageReader.surface)



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

        // 开启摄像头
        cameraManager.openCamera(cameraId, object : CameraDevice.StateCallback() {
            override fun onOpened(camera: CameraDevice) {
                Log.d(TAG, "camera open")
                cameraDevice = camera

//                创建会话
                camera.createCaptureSession(
                    SessionConfiguration(SessionConfiguration.SESSION_REGULAR,
                        listOf(cfg),
                        { it.run() },
                        object : CameraCaptureSession.StateCallback() {
                            override fun onConfigured(session: CameraCaptureSession) {
                                Log.i(TAG, "onConfigured: ")
                                // 创建捕获
                                val captureRequestBuild =
                                    session.device.createCaptureRequest(CameraDevice.TEMPLATE_STILL_CAPTURE)
                                        .apply {
                                            addTarget(imageReader.surface)
                                        }


                                // 开始捕获
                                session.setRepeatingRequest(captureRequestBuild.build(),
                                    null,
                                    cameraHandler)

                            }

                            override fun onConfigureFailed(session: CameraCaptureSession) {
                                Log.e(TAG, "onConfigureFailed: ")
                                session.close()
                            }

                        }))

            }

            override fun onDisconnected(camera: CameraDevice) {
                Log.d(TAG, "camera onDisconnected")
                camera.close()
                cameraDevice = null
            }

            override fun onError(camera: CameraDevice, error: Int) {
                Log.d(TAG, "camera onError")
                cameraDevice?.close()
                cameraDevice = null
            }

        }, cameraHandler)
    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun reConfiguration(callback: (session: CameraCaptureSession?) -> Unit) {
        _session?.close()
        cameraDevice?.let { createCaptureSession(it, callback) }
    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun resetRepeatingRequest(session: CameraCaptureSession) {
        // 创建捕获
        val captureRequestBuild =
            session.device.createCaptureRequest(CameraDevice.TEMPLATE_PREVIEW)
                .apply {
                    targets.forEach { surface -> addTarget(surface) }
                }.build()


        // 开始捕获
        session.setRepeatingRequest(captureRequestBuild,
            null,
            cameraHandler)


    }


    fun stopScan() {
        cameraDevice?.close()
    }

    fun release() {
        cameraThread.quitSafely()
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

        fun SAMPLE_YUV_420_888toNV21(image: Image): ByteArray {
            val nv21: ByteArray
            val yBuffer = image.planes[0].buffer
            val uBuffer = image.planes[1].buffer
            val vBuffer = image.planes[2].buffer
            val ySize = yBuffer.remaining()
            val uSize = uBuffer.remaining()
            val vSize = vBuffer.remaining()
            nv21 = ByteArray(ySize + uSize + vSize)

            //U and V are swapped
            yBuffer[nv21, 0, ySize]
            vBuffer[nv21, ySize, vSize]
            uBuffer[nv21, ySize + vSize, uSize]
            return nv21
        }
    }
}

fun Image.YUV_420_888toNV21(): ByteArray? {
    return if (format == ImageFormat.YUV_420_888) {
        val data = ByteArray(planes[0].buffer.capacity() * 3 / 2)
        val buff0Offset: Int = planes[0].buffer.capacity()

        planes[0].buffer.get(data, 0, buff0Offset)
        planes[2].buffer.get(data, buff0Offset, planes[2].buffer.capacity())
        data
    } else {
        null
    }
}

