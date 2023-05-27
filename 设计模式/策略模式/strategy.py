#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   strategy.py
@Time    :   2022/09/13 15:37:05
@Author  :   scliang 
@Desc    :   策略模式
'''

'''
定义一系列算法，把它们一个个封装起来，使它们可以相互替换；本模式使得算法可独立于使用它的客户而变化.
角色：
    抽象策略
    具体策略
    上下文
优点：
    1. 定义了一系列可重用的算法和行为.
    2. 消除了一些条件语句.
    3. 可以提供相同行为的不同实现.
缺点：
    client必须了解不同的策略
'''


# here put the import lib
from abc import abstractmethod, ABCMeta
from datetime import datetime

# 抽象策略
class Strategy(metaclass=ABCMeta):
    @abstractmethod
    def execute(self, data):
        pass

# 具体策略
class FastStrategy(Strategy):
    def execute(self, data):
        print("使用较快的策略处理%s" % data)

# 具体策略
class SlowStrategy(Strategy):
    def execute(self, data):
        print("使用较慢的策略处理%s" % data)

# 上下文
class Context:
    def __init__(self, strategy, data):
        self.data = data
        self.strategy = strategy
        # 可以定义用户不知道的东西
        self.date = datetime.now()

    def set_strategy(self, strategy):
        self.strategy = strategy

    def do_strategy(self):
        self.strategy.execute(self.data)

data = "Hello python"
# 使用较快的策略处理
fast_strategy = FastStrategy()
context = Context(fast_strategy, data)
context.do_strategy()
# 使用较慢的策略处理
slow_strategy = SlowStrategy()
context = Context(slow_strategy, data)
context.do_strategy()
"""
使用较快的策略处理Hello python
使用较慢的策略处理Hello python
"""