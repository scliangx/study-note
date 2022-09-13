#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   factory_method.py
@Time    :   2022/09/12 21:59:33
@Author  :   scliang 
@Desc    :   工厂方法模式
'''

'''
定义一个用于创建对象的接口（工厂接口），让子类决定实例化那一类产品.
优点： 
    1. 每个具体产品都对应一个具体工厂类，不需要修改工厂类代码.
    2. 隐藏了对象创建实现细节.
缺点：
    1. 每增加一个具体的产品类，都必须增加一个相应的具体工厂类. 
'''


# here put the import lib

from abc import ABCMeta,abstractmethod
from importlib.metadata import metadata


class Payment(metaclass=ABCMeta):
    # abstract class
    @abstractmethod
    def pay(self,money):
        pass


# 实现抽象类
class AliPay(Payment):
    def pay(self,money):
       print("支付宝支付%d" %money)

class WechatPay(Payment):
    def pay(self, money):
        print("微信支付%d元" % money)


# 创建一个工厂类接口
class PaymentFactory(metaclass=ABCMeta):
    @abstractmethod
    def create_payment(self,money):
        pass


# 阿里对象实现工厂类
class AlipayFactory(PaymentFactory):
    def create_payment(self):
        return AliPay()


# 微信对象实现工厂类
class WechatFectory(PaymentFactory):
    def create_payment(self):
        return WechatPay()

# client
pf = WechatFectory()
p = pf.create_payment()
p.pay(100)