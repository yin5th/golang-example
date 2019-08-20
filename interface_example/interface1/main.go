package main

import (
    "fmt"
)

/** 抽象层  **/
type Card interface {
    Display()
}

type Memory interface {
    Storage()
}

type CPU interface {
    Calculate()
}

type Computer struct {
    cpu CPU
    mem Memory
    card Card
}

func NewComputer(cpu CPU,mem Memory,card Card) *Computer {
    return &Computer{
        cpu:cpu,
        mem:mem,
        card:card,
    }
}

func (c *Computer) DoWork() {
    c.cpu.Calculate()
    c.mem.Storage()
    c.card.Display()
}

/** 实现层 **/
type Intel struct {
    //
}

type Kingston struct {
    //
}

type NVIDIA struct {
    //
}

func (i *Intel) Display(card Card) {
    fmt.Println("Intel 显卡运行中...")
}

func (i *Intel) Storage(memory Memory) {
    fmt.Println("Intel 内存运行中...")
}

func (i *Intel) Calculate(cpu CPU) {
    fmt.Println("Intel CPU运行中...")
}

func (k *Kingston) Storage(memory Memory) {
    fmt.Println("Kingston内存运行中...")
}

func (n *NVIDIA) Display(card Card) {
    fmt.Println("NVIDIA 显卡运行中...")
}

/** 业务逻辑层 **/

func main() {
    //intel系列电脑
    c1 := NewComputer(&)
}
