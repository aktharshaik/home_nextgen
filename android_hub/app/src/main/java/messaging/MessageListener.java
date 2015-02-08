package messaging;

import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.telephony.SmsMessage;
import android.util.Log;

import java.util.HashMap;
import java.util.Map;

/**
 * Created by ravitejareddy on 23/11/14.
 */
public class MessageListener extends BroadcastReceiver {

    private SharedPreferences preferences;
    public static final String TAG = MessageListener.class.getSimpleName();

    @Override
    public void onReceive(Context context, Intent intent) {
        // TODO Auto-generated method stub

        if (intent.getAction().equals("android.provider.Telephony.SMS_RECEIVED")) {
            Bundle bundle = intent.getExtras();           //---get the SMS message passed in---
            SmsMessage[] msgs = null;
            String msg_from = null;
            Map<String, String> msgMap = null;
            if (bundle != null) {
                //---retrieve the SMS message received---
                try {
                    Object[] pdus = (Object[]) bundle.get("pdus");
                    msgMap = new HashMap<String, String>(pdus.length);
                    msgs = new SmsMessage[pdus.length];
                    for (int i = 0; i < msgs.length; i++) {
                        msgs[i] = SmsMessage.createFromPdu((byte[]) pdus[i]);
                        msg_from = msgs[i].getOriginatingAddress();
                        if (!msgMap.containsKey(msg_from)) {
                            // Index with number doesn't exist
                            // Save string into associative array with sender number as index
                            msgMap.put(msgs[i].getOriginatingAddress(), msgs[i].getDisplayMessageBody());
                        } else {
                            // Number has been there, add content but consider that
                            // msg.get(originatinAddress) already contains sms:sndrNbr:previousparts of SMS,
                            // so just add the part of the current PDU
                            String previousparts = msgMap.get(msg_from);
                            String msgString = previousparts + msgs[i].getMessageBody();
                            msgMap.put(msg_from, msgString);
                        }
                    }
                    Log.d(TAG, msg_from +":"+ msgMap.get(msg_from));
                } catch (Exception e) {
//                            Log.d("Exception caught",e.getMessage());
                }
            }
        }
    }
}
