package com.example.ravitejareddy.home_nextgen;

import android.app.Service;
import android.content.Intent;
import android.os.IBinder;
import android.util.Log;

import java.io.IOException;
import java.net.ServerSocket;
import java.net.Socket;
import java.net.SocketException;

import messaging.MessageListener;

/**
 * Created by ravitejareddy on 15/11/14.
 */
public class UpdaterService extends Service {
    private static final String TAG = UpdaterService.class.getSimpleName();
    private Updater updater;
    private boolean isRunning = false;

    private ServerSocket serverSocket;
    public static final int PORT = 9999;
    private MessageListener messageListener;

    @Override
    public void onCreate() {
        super.onCreate();
        Log.d(TAG, "On Create");
        updater = new Updater();
        updater.start();
        messageListener = new MessageListener();

    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        Log.d(TAG, "On Destroy");
        if (serverSocket != null) {
            try {
                serverSocket.close();
            } catch (IOException e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }
        }
    }

    @Override
    public IBinder onBind(Intent intent) {
        return null;
    }


    class Updater extends Thread{
        static final long DELAY = 10000;

        Updater() {
            super("UPDATER");
        }

        @Override
        /* Do socket programming here */
        public void run() {
            int count = 0;
            String message = "";
            isRunning = true;
            Socket socket;
            super.run();

            /* Create a server socket in separate thread */
            try {
                serverSocket = new ServerSocket(PORT, 20);
            } catch (IOException e) {
                e.printStackTrace();
                Log.d(TAG, "FAILED TO CREATE SERVER SOCKET");
            }

            while (true){
                try {
                    Log.d(TAG, "LISTENING");
                    socket = serverSocket.accept();
                    Log.d(TAG, "ACCEPTED");
                    ++count;
                    ServerReplyThread serverReplyThread = new ServerReplyThread(socket, count);
                    serverReplyThread.start();
                } catch (SocketException e){
                    Log.d(TAG, "Couldn't connect");
                } catch (IOException e) {
                    e.printStackTrace();
                    Log.d(TAG, "IO Problem");

                }

            }
        }
    }


}
