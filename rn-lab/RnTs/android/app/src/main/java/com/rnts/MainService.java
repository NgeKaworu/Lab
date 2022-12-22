package com.rnts;

import android.app.IntentService;
import android.app.Notification;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.app.Service;
import android.content.Intent;
import android.os.Build;
import android.os.IBinder;
import android.os.Looper;
import android.util.Log;
import android.widget.Toast;

import androidx.annotation.Nullable;
import androidx.annotation.RequiresApi;

public class MainService extends Service {
    boolean isRun = false;

    @RequiresApi(api = Build.VERSION_CODES.O)
    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        Toast.makeText(this, "service starting", Toast.LENGTH_SHORT).show();
        Intent notificationIntent = new Intent(this, MainActivity.class);
        PendingIntent pendingIntent =
                PendingIntent.getActivity(this, 0, notificationIntent,
                        PendingIntent.FLAG_IMMUTABLE);


        isRun = true;
        new Thread(() -> {
            int count = 0;
            while (isRun) {
                Log.d("TAG", "Background Run ......");
                Notification notification =
                        new Notification.Builder(this, getString(R.string.channel_name))
                                .setContentTitle(getText(R.string.notification_title))
                                .setContentText(getText(R.string.notification_message) + " : " + count)
                                .setSmallIcon(R.drawable.ic_stat_name)
                                .setContentIntent(pendingIntent)
                                .setTicker(getText(R.string.ticker_text))
                                .build();
                startForeground(1, notification);

                NotificationManager notificationManager = getSystemService(NotificationManager.class);
                notificationManager.notify(1, notification);
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }).start();
        return super.onStartCommand(intent, flags, startId);
    }


    @Override
    public void onDestroy() {

        super.onDestroy();
        isRun = false;
    }

    @Nullable
    @Override
    public IBinder onBind(Intent intent) {
        return null;
    }
}
