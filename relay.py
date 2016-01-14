import time
import smbus
import sys

bus = smbus.SMBus(1)    # 0 = /dev/i2c-0 (port I2C0), 1 = /dev/i2c-1 (port I2C1)
data = int(sys.argv[1])
print ~data
#bus.write_byte_data(0x20, 0x06, ~data)
