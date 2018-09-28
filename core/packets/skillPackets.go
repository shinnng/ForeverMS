package packets

import (
	"github.com/sryanyuan/ForeverMS/core/maplepacket"
)

func SkillMelee(charID int32, skillID int32, targets, hits, display, animation byte, damages map[int32][]int32) maplepacket.Packet {
	p := maplepacket.NewPacket()
	p.WriteByte(maplepacket.SendChannelPlayerUseStandardSkill)
	p.WriteInt32(charID)
	p.WriteByte(byte(targets*0x10) + hits)
	p.WriteBool(bool(skillID != 0))
	if skillID != 0 {
		p.WriteInt32(skillID)
	}
	p.WriteByte(display)
	p.WriteByte(animation)

	p.WriteByte(0)  // mastery
	p.WriteInt32(0) // starID?

	for k, v := range damages {
		p.WriteInt32(k)
		p.WriteByte(0x6)
		// if meos explosion add, another byte for something
		for _, dmg := range v {
			p.WriteInt32(dmg)
		}
	}

	return p
}

func SkillRanged(charID, skillID, objID int32, targets, hits, display, animation byte, damages map[int32][]int32) maplepacket.Packet {
	p := maplepacket.NewPacket()
	p.WriteByte(maplepacket.SendChannelPlayerUseRangedSkill)
	p.WriteInt32(charID)
	p.WriteByte(targets*0x10 + hits)
	p.WriteBool(bool(skillID != 0))
	if skillID != 0 {
		p.WriteInt32(skillID)
	}
	p.WriteByte(display)
	p.WriteByte(animation)

	p.WriteByte(0)      // mastery
	p.WriteInt32(objID) // starID?

	for k, v := range damages {
		p.WriteInt32(k)
		p.WriteByte(0x6)
		for _, dmg := range v {
			p.WriteInt32(dmg)
		}
	}

	return p
}

func SkillMagic(charID int32, skillID int32, targets, hits, display, animation byte, damages map[int32][]int32) maplepacket.Packet {
	p := maplepacket.NewPacket()
	p.WriteByte(maplepacket.SendChannelPlayerUseMagicSkill)
	p.WriteInt32(charID)
	p.WriteByte(targets*0x10 + hits)
	p.WriteBool(bool(skillID != 0))
	if skillID != 0 {
		p.WriteInt32(skillID)
	}
	p.WriteByte(display)
	p.WriteByte(animation)

	p.WriteByte(0)  // mastery
	p.WriteInt32(0) // starID?

	for k, v := range damages {
		p.WriteInt32(k)
		p.WriteByte(0x6)
		for _, dmg := range v {
			p.WriteInt32(dmg)
		}
	}

	return p
}

func SkillAnimation(charID int32, skillID int32, level byte) maplepacket.Packet {
	p := maplepacket.NewPacket()
	p.WriteByte(maplepacket.SendChannelPlayerAnimation)
	p.WriteInt32(charID)
	p.WriteByte(0x01)
	p.WriteInt32(skillID)
	p.WriteByte(level)

	return p
}

func SkillGmHide(isHidden bool) maplepacket.Packet {
	p := maplepacket.NewPacket()
	p.WriteByte(maplepacket.SendChannelEmployee)
	p.WriteByte(0x0F)
	p.WriteBool(isHidden)

	return p
}
