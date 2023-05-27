#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   bridge.py
@Time    :   2022/09/13 12:09:53
@Author  :   scliang 
@Desc    :   桥模式
'''

'''
将一个事物的两个维度分离，使其都可以独立的变化.
角色：
    抽象
    细化抽象
    实现者
    具体实现者
优点：
    1. 抽象和实现相分离
    2. 扩展能力优秀
'''

# here put the import lib

from abc import ABCMeta, abstractmethod

class Shape(metaclass=ABCMeta):
    def __init__(self,color) -> None:
        self.color = color

    @abstractmethod
    def draw(self):
        pass

class Color(metaclass=ABCMeta):
    @abstractmethod
    def paint(self,shape):
        pass


class Circle(Shape):
    name = "圆形"
    def draw(self):
        self.color.paint(self)

class Rectangle(Shape):
    name = "长方形"
    def draw(self):
        self.color.paint(self)

class Red(Color):
    def paint(self, shape):
        print("红色的%s"%shape.name)

class Green(Color):
    def paint(self, shape):
        print("绿色的%s"%shape.name)

r = Rectangle(Red())
r.draw()