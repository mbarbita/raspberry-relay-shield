import time
import smbus
import sys

bus = smbus.SMBus(1)    # 0 = /dev/i2c-0 (port I2C0), 1 = /dev/i2c-1 (port I2C1)

#bus.write_byte_data(0x20, 0x06, 0xff)
#bus.write_byte_data(0x20, 0x06, int(sys.argv[1])
