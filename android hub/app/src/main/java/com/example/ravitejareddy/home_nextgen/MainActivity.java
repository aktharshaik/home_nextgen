package com.example.ravitejareddy.home_nextgen;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.os.Debug;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.view.WindowManager;
import android.widget.Button;
import android.widget.TextView;

import java.net.ServerSocket;


public class MainActivity extends Activity implements View.OnClickListener{

    TextView textView;
    Button startButton, stopButton;


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        getWindow().addFlags(WindowManager.LayoutParams.FLAG_KEEP_SCREEN_ON);

        textView = (TextView) findViewById(R.id.text_view);
        startButton = (Button)findViewById(R.id.itemServiceStart);
        stopButton = (Button)findViewById(R.id.itemServiceStop);
        startButton.setOnClickListener(this);
        stopButton.setOnClickListener(this);
        //Debug.startMethodTracing();
    }


    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;
    }

    @Override
    protected void onStop() {
        super.onStop();
        //Debug.stopMethodTracing();
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }

    @Override
    public void onClick(View view) {
        switch (view.getId()){
            case R.id.itemServiceStart:
                textView.setText("START SERVICE");
                Log.d("RAVI", "TEST");
                startService(new Intent(this, UpdaterService.class));
                break;
            case R.id.itemServiceStop:
                textView.setText("STOP SERVICE");
                stopService(new Intent(this, UpdaterService.class));
                break;
        }
    }
}
