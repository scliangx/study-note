#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   interacfe.py
@Time    :   2022/09/12 21:21:05
@Author  :   scliang 
@Desc    :   None
'''

# here put the import lib
# 引入抽象类
from abc import ABCMeta, abstractmethod
import re

class Payment(metaclass=ABCMeta):
    # abstract class
    @abstractmethod
    def pay(self,money):
        pass


# 实现抽象类
class AliPay(Payment):
    def pay(self,moey):
        pass

class WechatPay(Payment):
    def pay(self, money):
        pass
