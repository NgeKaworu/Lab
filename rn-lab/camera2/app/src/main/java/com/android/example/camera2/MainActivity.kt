package com.android.example.camera2

import android.Manifest
import android.content.Context
import android.content.pm.PackageManager
import android.graphics.BitmapFactory
import android.graphics.ImageFormat
import android.graphics.Rect
import android.graphics.YuvImage
import android.hardware.camera2.*
import android.media.ImageReader
import android.os.Build
import android.os.Bundle
import android.os.Handler
import android.os.HandlerThread
import android.util.Log
import androidx.annotation.RequiresApi
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import androidx.core.content.ContextCompat
import com.android.example.camera2.databinding.ActivityMainBinding
import java.io.ByteArrayOutputStream


class MainActivity : AppCompatActivity() {
    private var prepared = false
    private lateinit var viewBinding: ActivityMainBinding;
    private lateinit var cameraManager: CameraManager
    private lateinit var cameraId: String
    private var cameraDevice: CameraDevice? = null
    private var imageReader: ImageReader? = null

    /** [HandlerThread] where all camera operations run */
    private val cameraThread = HandlerThread("CameraThread").apply { start() }
    private val cameraHandler = Handler(cameraThread.looper)

    private val faceScanner by lazy {
        FaceScanner(this)
    }


    @RequiresApi(Build.VERSION_CODES.P)
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        viewBinding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(viewBinding.root)
        if (!allPermissionsGranted(this)) {
            ActivityCompat.requestPermissions(this,
                REQUIRED_PERMISSIONS,
                REQUEST_CODE_PERMISSIONS)
        }

        viewBinding.cameraToggle.setOnClickListener {
            faceScanner.toggleCamera {
                Log.i(TAG, "cameraToggle click: $it")
            }
        }


    }

    @RequiresApi(Build.VERSION_CODES.P)
    fun f1() {
        faceScanner.startScan(640, 480) {
            val frame = it.acquireNextImage();
            val nv21 = FaceScanner.YUV_420_888toNV21(frame)

//                val p = frame.planes
//                val buffer = p[0].buffer
//                val bytes = ByteArray(buffer.remaining())
//                buffer.get(bytes)


            val img = YuvImage(nv21, ImageFormat.NV21, frame.width, frame.height, null)
            val stream = ByteArrayOutputStream()
            img.compressToJpeg(Rect(0, 0, frame.width, frame.height), 80, stream)


            val bitmap = BitmapFactory.decodeByteArray(stream.toByteArray(), 0, stream.size())
            stream.close()

            runOnUiThread {
                viewBinding.imageView.setImageBitmap(bitmap)
            }


//                Log.i(TAG, "onCreate: $nv21")
            frame?.close()
        }

    }

    fun f2() {
        faceScanner.stopScan()
    }


    companion object {
        private const val TAG = "Camera2"
        private const val REQUEST_CODE_PERMISSIONS = 10
        private fun allPermissionsGranted(context: Context) = REQUIRED_PERMISSIONS.all {
            ContextCompat.checkSelfPermission(context, it) == PackageManager.PERMISSION_GRANTED
        }

        val REQUIRED_PERMISSIONS = mutableListOf(
            Manifest.permission.CAMERA,
            Manifest.permission.RECORD_AUDIO,
            Manifest.permission.WRITE_EXTERNAL_STORAGE,
            Manifest.permission.READ_EXTERNAL_STORAGE,
        ).toTypedArray();
    }

}


