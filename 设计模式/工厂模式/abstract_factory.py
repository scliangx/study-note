#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   factory.py
@Time    :   2022/09/12 21:59:33
@Author  :   scliang 
@Desc    :   抽象工厂模式
'''

'''
定义一个用于创建对象的接口（工厂接口），让工厂子类来创建一系列相关或相互依赖的对象.
相比工厂方法模式，抽象工厂模式中每个具体工厂都生产一套产品.
优点： 
    1. 将客户端与类的具体实现相互分离.
    2. 每个工厂创建了一个完整的产品系列，使得易于交换产品系列.
    3. 有利于产品的一致性(即产品之间的约束关系)
缺点：
    1. 难以支持新品种的(抽象)产品. 
'''


# here put the import lib
from abc import ABCMeta,abstractmethod


# 定义抽象产品
class PhoneShell(metaclass=ABCMeta):
    @abstractmethod
    def show_shell(self):
        pass

class CPU(metaclass=ABCMeta):
    def show_cpu(self):
        pass

class OS(metaclass=ABCMeta):
    def show_os(self):
        pass


# 抽象工厂
class PhoneFectory(metaclass=ABCMeta):
    @abstractmethod
    def make_shell(self):
        pass

    @abstractmethod
    def make_cpu(self):
        pass

    @abstractmethod
    def make_os(self):
        pass

# 具体产品
class SmallShell(PhoneShell):
    def show_shell(self):
        print("普通的小手机壳.")

class BigShell(PhoneShell):
    def show_shell(self):
        print("普通的大手机壳.")

class AppleShell(PhoneShell):
    def show_shell(self):
        print("apply 手机壳.")

class SnapDragon(CPU):
    def show_cpu(self):
        print("骁龙CPU.")


class MediaTekCPU(CPU):
    def show_cpu(self):
        print("联发科CPU.")

class ApplyCPU(CPU):
    def show_cpu(self):
        print("apply CPU.")

class Android(OS):
    def show_os(self):
        print("Android OS.")

class IOS(OS):
    def show_os(self):
        print("IOS OS.")


# 具体的工厂
class MiFactory(PhoneFectory):
    '''
    创建了shell, cpu, os三个完整的产品
    '''

    def make_cpu(self):
        return SnapDragon()

    def make_os(self):
        return Android()

    def make_shell(self):
        return BigShell()

class HuaweiFactory(PhoneFectory):
    '''
    创建了shell, cpu, os三个完整的产品
    '''

    def make_shell(self):
        return MediaTekCPU()

    def make_os(self):
        return Android()
    
    def make_cpu(self):
        return SmallShell()

class IPhoneFactory(PhoneFectory):
    '''
    创建了shell, cpu, os三个完整的产品
    '''

    def make_shell(self):
        return AppleShell()

    def make_os(self):
        return IOS()

    def make_cpu(self):
        return ApplyCPU()


#  client
class Phone():
    def __init__(self,cpu,os,shell) -> None:
        self.cpu = cpu
        self.shell = shell
        self.os = os

    def show_info(self):
        print("手机信息: ")
        self.cpu.show_cpu()
        self.os.show_os()
        self.shell.show_shell()

def make_phone(factory):
    print("factory type: ",type(factory))
    cpu = factory.make_cpu()
    os = factory.make_os()
    shell = factory.make_shell()
    return Phone(cpu,os,shell)

p1 = make_phone(MiFactory())
p1.show_info()
p2 = make_phone(IPhoneFactory())
p2.show_info()

'''
PS F:\GoProjects\src\MyPractiseNotes\设计模式\工厂模式> python .\abstract_factory.py
factory type:  <class '__main__.MiFactory'>
手机信息: 
骁龙CPU.
Android OS.
普通的大手机壳.
factory type:  <class '__main__.IPhoneFactory'>
手机信息:
apply CPU.
IOS OS.
apply 手机壳.
'''
