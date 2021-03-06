package model

import (
	"common/utils"
	"fmt"
	"encoding/json"
)

// 机器监控和实例监控都会产生Event，共用这么一个struct
type Event struct {
	Id          string            `json:"id"`
	Strategy    *Strategy         `json:"strategy"`
	Expression  *Expression       `json:"expression"`
	Status      string            `json:"status"` // OK or PROBLEM
	Endpoint    string            `json:"endpoint"`
	LeftValue   float64           `json:"leftValue"`
	CurrentStep int               `json:"currentStep"`
	EventTime   int64             `json:"eventTime"`
	PushedTags  map[string]string `json:"pushedTags"`
}

func (this *Event) FormattedTime() string {
	return utils.UnixTsFormat(this.EventTime)
}

func (this *Event) String() string {
	return fmt.Sprintf(
		"<Endpoint:%s, Status:%s, Strategy:%v, Expression:%v, LeftValue:%s, CurrentStep:%d, PushedTags:%v, TS:%s>",
		this.Endpoint,
		this.Status,
		this.Strategy,
		this.Expression,
		utils.ReadableFloat(this.LeftValue),
		this.CurrentStep,
		this.PushedTags,
		this.FormattedTime(),
	)
}

func (this *Event) ExpressionId() int {
	if this.Expression != nil {
		return this.Expression.Id
	}

	return 0
}

func (this *Event) TagsString() string {
	bs, err := json.Marshal(this.PushedTags)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", string(bs))
}

func (this *Event) StrategyId() int {
	if this.Strategy != nil {
		return this.Strategy.Id
	}

	return 0
}

func (this *Event) TplId() int {
	if this.Strategy != nil {
		return this.Strategy.Tpl.Id
	}

	return 0
}

func (this *Event) ActionId() int {
	if this.Expression != nil {
		return this.Expression.ActionId
	}

	return this.Strategy.Tpl.ActionId
}

func (this *Event) Priority() int {
	if this.Strategy != nil {
		return this.Strategy.Priority
	}
	return this.Expression.Priority
}

func (this *Event) AlarmLevel() string {
	switch this.Priority() {
	default:
		return "未知告警"
	case 0:
		return "致命告警"
	case 1:
		return "高等级告警"
	case 2, 3, 4:
		return "中等级告警"
	case 5, 6:
		return "低等级告警"
	}
}

func (this *Event) StatusString() string {
	switch this.Status {
	default:
		return "未知状态"
	case "OK":
		return "已恢复"
	case "PROBLEM":
		return "未恢复"
	}
}

func (this *Event) Note() string {
	if this.Strategy != nil {
		return this.Strategy.Note
	}
	return this.Expression.Note
}

func (this *Event) Metric() string {
	if this.Strategy != nil {
		return this.Strategy.Metric
	}
	return this.Expression.Metric
}

func (this *Event) RightValue() float64 {
	if this.Strategy != nil {
		return this.Strategy.RightValue
	}
	return this.Expression.RightValue
}

func (this *Event) Operator() string {
	if this.Strategy != nil {
		return this.Strategy.Operator
	}
	return this.Expression.Operator
}

func (this *Event) Func() string {
	if this.Strategy != nil {
		return this.Strategy.Func
	}
	return this.Expression.Func
}

func (this *Event) MaxStep() int {
	if this.Strategy != nil {
		return this.Strategy.MaxStep
	}
	return this.Expression.MaxStep
}

func (this *Event) Counter() string {
	return fmt.Sprintf("%s/%s %s", this.Endpoint, this.Metric(), utils.SortedTags(this.PushedTags))
}
