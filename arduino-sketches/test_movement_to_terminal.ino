#include <Servo.h>

Servo servo1;

void setup() {
    Serial.begin(9600);
    servo1.attach(10);
    servo1.write(80);
}

void loop() {
    if (Serial.available() > 0) {
      delay(2000);
      servo1.write(0);
    
      delay(2000);
      servo1.write(80);
      Serial.read();
    }
}