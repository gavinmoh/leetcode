package main

import "sort"

type Packet struct {
	Source      int
	Destination int
	Timestamp   int
}

type Router struct {
	Packets            []Packet
	Indices            map[Packet]bool
	DestinationPackets map[int][]Packet
	MemoryLimit        int
}

func Constructor(memoryLimit int) Router {
	return Router{
		Packets:            []Packet{},
		Indices:            map[Packet]bool{},
		DestinationPackets: map[int][]Packet{},
		MemoryLimit:        memoryLimit,
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	packet := Packet{
		Source:      source,
		Destination: destination,
		Timestamp:   timestamp,
	}

	if _, ok := this.Indices[packet]; ok {
		return false
	}

	this.Indices[packet] = true
	this.Packets = append(this.Packets, packet)
	this.DestinationPackets[destination] = append(this.DestinationPackets[destination], packet)

	if len(this.Packets) > this.MemoryLimit {
		oldest := this.Packets[0]
		delete(this.Indices, oldest)
		this.DestinationPackets[oldest.Destination] = this.DestinationPackets[oldest.Destination][1:]
		this.Packets = this.Packets[1:]
	}

	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.Packets) == 0 {
		return []int{}
	}

	packet := this.Packets[0]
	delete(this.Indices, packet)
	this.DestinationPackets[packet.Destination] = this.DestinationPackets[packet.Destination][1:]
	this.Packets = this.Packets[1:]

	return []int{packet.Source, packet.Destination, packet.Timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	n := len(this.DestinationPackets[destination])
	left := sort.Search(n, func(i int) bool {
		return this.DestinationPackets[destination][i].Timestamp >= startTime
	})
	right := sort.Search(n, func(i int) bool {
		return this.DestinationPackets[destination][i].Timestamp > endTime
	})

	return right - left
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
