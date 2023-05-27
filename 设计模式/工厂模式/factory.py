#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   factory.py
@Time    :   2022/09/12 21:59:33
@Author  :   scliang 
@Desc    :   简单工厂模式
'''

'''
不直接向客户端暴露创建对象的细节，而是通过一个工厂类创建来负责创建对象的实例.
优点： 
    1. 隐藏了对象创建的细节
    2. 客户端不需要修改代码
缺点：
    1. 违反了单一职责原则，将创建逻辑集中到一个工厂类中 
    2. 当添加新产品时，需要修改工厂类代码，违反了闭开原则
'''


# here put the import lib

from abc import ABCMeta,abstractmethod


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

# 创建一个简单工厂类
# 根据需要返回不同的类对象
class PaymentFactory():
    def create_payment(self,method):
        if method == "alipay":
            return AliPay()
        elif method == "wechat":
            return WechatPay()
        else:
            raise TypeError("No such payment name %s" % method)

