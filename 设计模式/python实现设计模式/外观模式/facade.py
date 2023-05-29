#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   facade.py
@Time    :   2022/09/13 13:02:51
@Author  :   scliang 
@Desc    :   外观模式
'''

'''
为子系统的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这一接口使得这一子系统更加容易实现.
角色：
    外观
    子系统类
'''

# here put the import lib
# 子系统类
class CPU:
    def run(self):
        print('CPU start running...')

    def stop(self):
        print('CPU stop running...')

# 子系统类
class Disk:
    def run(self):
        print('Disk start running...')

    def stop(self):
        print('Disk stop running...')

# 子系统类
class Memory:
    def run(self):
        print('Memory start running...')

    def stop(self):
        print('Memory stop running...')

# 外观
class Computer():
    def __init__(self):
        self.CPU = CPU()
        self.Disc = Disk()
        self.Member = Memory()

    def run(self):
        self.CPU.run()
        self.Disc.run()
        self.Member.run()

    def stop(self):
        self.CPU.stop()
        self.Disc.stop()
        self.Member.stop()

# client
c = Computer()
c.run()
c.stop()