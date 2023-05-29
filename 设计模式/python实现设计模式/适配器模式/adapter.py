#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   adapter.py
@Time    :   2022/09/13 11:50:05
@Author  :   scliang 
@Desc    :   适配器模式
'''

'''
将一个类的接口转换为客户希望的另一个接口，适配器模式使得原本由于接口不兼容而不能一起工作的类可以一起工作.
想使用一些已经存在的子类，但不可能对每一个都进行子类化以匹配它们的接口，对象适配器可以适配它们父类的接口.
角色:
    目标接口
    待适配的类
    适配器
实现方式:
    1. 类适配器(多继承)
    2. 对象适配器(使用组合)
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

# 接口没有统一
class BankPay():
    def cost(self,money):
        print("银行卡支付%d元." % money)

class ApplyPay():
    def cost(self,money):
        print("Apply卡支付%d元." % money)

# 使用多继承创建适配器统一接口
# 将bankpay转换成payment统一的接口
class NewBanckPay(Payment,BankPay):
    def pay(self, money):
        self.cost(money)


# 使用组合的方式
class PaymentAdapter():
    def __init__(self,payment) -> None:
        self.payment = payment

    def pay(self,money):
        self.payment.cost(money)


# 传入一个适配类
adapter = PaymentAdapter(ApplyPay())
adapter.pay(1000)