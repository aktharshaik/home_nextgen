/*
The sketch demonstrates how to increase or decrease the Bluetooth Low Energy 4
Advertisement transmission interval.

Faster iterval = higher power consumption = lower connection latency
*/


#include <RFduinoBLE.h>

// pin 3 on the RGB shield is the green led
int led = 3;

// interval between advertisement transmissions ms (range is 20ms to 10.24s) - default 20ms
int duration = 10000;  // 675 ms between advertisement transmissions

void setup() {
  // led used to indicate that the RFduino is advertising
  pinMode(led, OUTPUT);

  // this is the data we want to appear in the advertisement
  // (if the deviceName and advertisementData are too long to fix into the 31 byte
  // ble advertisement packet, then the advertisementData is truncated first down to
  // a single byte, then it will truncate the deviceName)
  //RFduinoBLE.advertisementData = "Hello World\n";
  
  // change the advertisement interval
  //RFduinoBLE.advertisementInterval = interval;

  // start the BLE stack
  //RFduinoBLE.begin();
  //RFduinoBLE.connectable = false;
  Serial.begin(9600);
  Serial.println("Serial Begin");
}

void advertise(const char *data, uint32_t ms)
{
  Serial.println(data);
  // this is the data we want to appear in the advertisement
  // (if the deviceName and advertisementData are too long to fix into the 31 byte
  // ble advertisement packet, then the advertisementData is truncated first down to
  // a single byte, then it will truncate the deviceName)
  RFduinoBLE.advertisementData = data;
  
  // start the BLE stack
  RFduinoBLE.begin();
  
  // advertise for ms milliseconds
  RFduino_ULPDelay(ms);
  
  // stop the BLE stack
  RFduinoBLE.end();
}

void loop() {
  // advertise "111" for indicated time
  advertise("HELLO", duration);
  
  // adverise "222" for indicated time
  advertise("WORLD", duration);
}
void RFduinoBLE_onConnect()
{
	Serial.println("RFduino BLE connection successful");
}

void RFduinoBLE_onDisconnect()
{
	Serial.println("RFduino BLE disconnected");
}

void RFduinoBLE_onReceive(char *data, int len)
{
	// If the first byte is 0x01 / on / true
	Serial.println("Received data over BLE");
	if (data[0])
	{
		Serial.println("Turn RFduino Blue LED On");
	}
	else
	{
		Serial.println("Turn RFduino Blue LED Off");
	}
}

void RFduinoBLE_onAdvertisement(bool start)
{
  // turn the green led on if we start advertisement, and turn it
  // off if we stop advertisement
  Serial.println("Advertisement");

}
