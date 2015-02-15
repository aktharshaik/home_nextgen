package com.example.ravitejareddy.home_nextgen;

import android.net.LocalSocket;
import android.util.Log;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintStream;
import java.net.Socket;
import java.net.SocketException;

/**
 * Created by ravitejareddy on 15/11/14.
 */
public class ServerReplyThread extends Thread {
    private Socket hostThreadSocket;
    public static final String TAG = "HOME_NXTGEN:" + ServerReplyThread.class.getSimpleName();
    int cnt;
    String message = "";

    ServerReplyThread(Socket socket, int c) {
        hostThreadSocket = socket;
        cnt = c;
    }

    public void run() {
        OutputStream outputStream;
        String msgReply = "Hello from Android, you are #" + cnt;

        //message = "#" + cnt + " from " + hostThreadSocket.getInetAddress() + ":" + hostThreadSocket.getPort() + "\n";
        Log.d(TAG, message);

        try {
            BufferedReader input = new BufferedReader(new InputStreamReader(hostThreadSocket.getInputStream()));
            String string = "";
            outputStream = hostThreadSocket.getOutputStream();
            PrintStream printStream = new PrintStream(outputStream);
            //printStream.print(msgReply);
            while(true){
                Log.d(TAG, "Waiting for input");
                string = input.readLine();
                Log.d(TAG, "Writing to client");
                printStream.print(string+"\n");
                Log.d(TAG, "Written");
            }
        } catch (SocketException e) {
            Log.d(TAG, "SOMETHING WENT WRONG");
        } catch (IOException e) {
            e.printStackTrace();
        }


    }
}
