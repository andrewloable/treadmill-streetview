int analogPin = 3;     
int val = 0;         
unsigned long pulses = 0L;
int previous = 0;

// modify variables as needed 
const int ANALOG_HIGH = 200;
const int PULSES_ONE_LOOP = 10000;
// modify varialbes as needed


void setup()
{
  Serial.begin(9600);              //  setup serial
  //pinMode(LED_BUILTIN, OUTPUT);
}

void loop()
{
  val = analogRead(analogPin);     // read the input pin
  if (val > ANALOG_HIGH) {
    checkPulse(1);
    //digitalWrite(LED_BUILTIN, HIGH);
  } else {
    checkPulse(0);
    //digitalWrite(LED_BUILTIN, LOW);
  }
  //Serial.println(pulses);
  if (pulses >= PULSES_ONE_LOOP){
    // send one loop message to serial
    Serial.print(pulses);
    Serial.print("ONE LOOP-");
    // reset pulses
    pulses = 0;
  }
}

// Check if rising or falling signal
void checkPulse(int input){
  if (previous != input){
    previous = input;
    pulses += 1;
  }
}

