#include <WiFiClient.h>
#include <ArduinoJson.h>
#include <ESP8266WiFi.h>
#include <ESP8266httpUpdate.h>
#include <ESP8266HTTPClient.h>

#define CurrentVersion 1.00  // Current Firmware version

WiFiClient client;
HTTPClient http;

unsigned long interval = 5000;  // Time period to check update in ms
unsigned long previousMillis = 0;

const char* ssid = "WiFi SSID";          // SSID (Access Point name)
const char* password = "WiFi PASSWORD";  // WiFi Password

void setup() {
  Serial.begin(115200);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    Serial.println("Connecting...");
  }
}

void update_started() {
  Serial.println("HTTP update process started");
}

void update_finished() {
  Serial.println("HTTP update process finished");
}

void update_progress(int cur, int total) {
  Serial.printf("HTTP update process at %d of %d bytes...\n", cur, total);
}

void update_error(int err) {
  Serial.printf("HTTP update fatal error code %d\n", err);
}

void otaUpdate(String url) {
  WiFiClient client;
  ESPhttpUpdate.setLedPin(LED_BUILTIN, LOW);
  ESPhttpUpdate.onStart(update_started);
  ESPhttpUpdate.onEnd(update_finished);
  ESPhttpUpdate.onProgress(update_progress);
  ESPhttpUpdate.onError(update_error);
  t_httpUpdate_return ret = ESPhttpUpdate.update(client, url);
  switch (ret) {
    case HTTP_UPDATE_FAILED:
      Serial.printf("HTTP_UPDATE_FAILD Error (%d): %s\n", ESPhttpUpdate.getLastError(), ESPhttpUpdate.getLastErrorString().c_str());
      break;
    case HTTP_UPDATE_NO_UPDATES:
      Serial.println("HTTP_UPDATE_NO_UPDATES");
      break;
    case HTTP_UPDATE_OK:
      Serial.println("HTTP_UPDATE_OK");
      break;
  }
}

void check_for_updates() {
  HTTPClient http;
  http.begin(client, "http://your-hostname-or-server-IP-address/updates/firmware.json");  // path to your hostname or server IP address where the firmware.json placed
  int httpCode = http.GET();
  String payload = "";
  if (httpCode > 0) {
    payload = http.getString();
  }
  http.end();  //Close connection
  String json = "";
  DynamicJsonDocument doc(1024);
  json = payload.c_str();
  DeserializationError error = deserializeJson(doc, json);
  if (error) {
  }
  if (doc["ver"] > CurrentVersion && doc["ver"] != CurrentVersion) {
    otaUpdate(doc["link"]);
  } else {
  }
}

void loop() {
  unsigned long currentMillis = millis();
  if (currentMillis - previousMillis >= interval) {
    check_for_updates();  //Check for update
    previousMillis = currentMillis;
  }
}