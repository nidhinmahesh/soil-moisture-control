# soilMoistureControl

#### ● Server fetch and analyse the temperature, humidity and

#### rain prediction data at the location at particular time.

#### ● Now cloud will decide a particular time for water launch.

#### ● MCU will match the time with clock and time from cloud

#### ● so at that particular time it will enable relay. which will

#### switch on the solenoid valve to control the water supply.


### BLOCK DIAGRAM & Temperature Sensor


## Soil Moisture Sensor FC-

#### 1. PARAMETERS: Volumetric content inside the soil

#### 2. HOW IT WORKS?

##### 1. The soil moisture sensor consists of two probes, to measure the volumetric content of water.

##### a. allow current to pass through the soil and then get the resistance value to measure

##### moisture value

##### 2. When there is more water

##### a. the soil will conduct more electricity i.e there will be less resistance. So, the moisture level

##### will be higher.

##### 3. when there will be less water

##### a. then the soil will conduct less electricity which means that there will be more resistance. So,

##### moisture level will be less.

##### 4. Can work in both the modes: Analog and Digital mode

##### .


##### 5. When the moisture in the soil is less than the set threshold, the output remains

##### low.

##### 6. The digital output is connected to a microcontroller to sense the moisture level.

```
Input Voltage 3.3 – 5V
```
```
Output Voltage 0 – 4.2V
```
```
Input Current 35mA
```
```
Output Signal Both Analog and Digital
```

### HUMIDITY AND TEMPERATURE SENSOR

#### 1. PARAMETERS: Atmospheric temperature and Relative

#### humidity


### DHT11 Digital Relative Humidity & Temp Sensor Module:

### How it Works?

###### ● 3 Major Parts:

###### 1. A NTC temperature sensor (or thermistor)

###### 2. A humidity sensing component

###### 3. An IC on the back side of the sensor


###### 1. FOR MEASURING TEMPERATURE:

###### ● uses a NTC temperature sensor or a thermistor

```
● thermistor: is a variable resistor that changes its
resistance with change of the temperature
● made by sintering of semiconductive materials
such as ceramics or polymers
● to provide larger changes in the resistance with
just small changes in temperature
● “NTC” means “ Negative Temperature
Coefficient ”
● means that the resistance decreases with
increase of the temperature
```
```
“NTC” : “ Negative Temperature Coefficient ”
```

###### 2. FOR MEASURING HUMIDITY:

```
● uses the humidity sensing component
● has two electrodes
● moisture holding substrate between
electrodes
● as the humidity changes, the conductivity
of the substrate changes
● so the resistance between these
electrodes changes
● this change in resistance is measured and
processed by the IC
● which makes it ready to be read by a
microcontroller
```

## Circuit Connection


## Solenoid Valve ● Valve senses that certain^

##### quantity of the flow of the

##### fluid is required

##### ● It allows the current to pass

##### through the solenoid valve.

##### ● Due to this, the valve gets

##### energized and the magnetic

##### field is generated

##### ● Which triggers the movement

##### of the plunger against the

##### action of the spring.

##### ● Plunger moves in upwards

##### direction, which allows the

##### opening of the orifice.

##### ● At this instant the flow of the

##### fluid is allowed from the inlet

##### port to the outlet port.


## Cloud do the forecast


## FIELD TESTING PLAN

#### There are two ways of measuring soil moisture

#### namely

#### ● Direct Inspection-Through physical means

#### ● Meters and Sensors-Through ESP8266 and

#### Arduino.


#### So our ultimate plan is to grow a chilli plant on the rooftop

#### using the soil moisture sensors.

#### Temperature and humidity are the two major factors involved

#### in the growth of chilli.

#### The maximum temperature should be 36 degree C and the

#### plants should enjoy high humidity.
