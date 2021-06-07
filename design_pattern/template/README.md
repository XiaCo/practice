# 模版模式

代码层面，一个模版接口，多个其子类

---

横向拓展。

过程一致，步骤细节由各个子类自己实现

如果子类有相同的部分，还可以实现一个base类，写一些通用逻辑

再说说和建造者模式的一些联想，从抽象层面来说，其实和建造者模式没有区别。

建造者模式只是比模版模式**多了个获取最终结果**的方法

它的*Template类*相当于创建者模式的*Director类*，都是用于对过程的管理与描述