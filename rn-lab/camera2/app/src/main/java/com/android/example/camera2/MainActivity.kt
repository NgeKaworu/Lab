package com.android.example.camera2

import android.Manifest
import android.content.Context
import android.content.pm.PackageManager
import android.graphics.*
import android.hardware.camera2.*
import android.os.*
import android.util.Log
import android.view.Surface
import androidx.annotation.RequiresApi
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import androidx.core.content.ContextCompat
import com.android.example.camera2.databinding.ActivityMainBinding


class MainActivity : AppCompatActivity() {
    private lateinit var viewBinding: ActivityMainBinding;


    private val faceScanner by lazy {
        Camera2Util(this)
    }


    @RequiresApi(Build.VERSION_CODES.P)
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        viewBinding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(viewBinding.root)
        if (!allPermissionsGranted(this)) {
            ActivityCompat.requestPermissions(this, REQUIRED_PERMISSIONS, REQUEST_CODE_PERMISSIONS)
        }


        viewBinding.sessionToggle.setOnClickListener {
            if (faceScanner.isPrepared()) {
                faceScanner.closeCamera()
            } else {
                faceScanner.prepare {}
            }
        }

        viewBinding.previewToggle.setOnClickListener {
            if (faceScanner.inPreviewing()) {
                faceScanner.stopPreview()
            } else {
                faceScanner.startPreview(Surface(viewBinding.textureView.surfaceTexture))
            }
        }

        viewBinding.scanToggle.setOnClickListener {
            if (faceScanner.inScanning()) {
                faceScanner.stopScan()
            } else {

                faceScanner.startScan {
                    val frame = it.acquireNextImage();
                    val nv21 = Camera2Util.YUV_420_888toNV21(frame)

                    Log.i(TAG, "onCreate: $nv21")
                    frame?.close()
                }

            }
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


