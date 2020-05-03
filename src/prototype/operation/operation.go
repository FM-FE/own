package operation

import "time"

type Tree struct {
	LeftNode  *Operation `json:"left_node"`
	Operation Operation  `json:"Operation"`
	RightNode *Operation `json:"right_node"`
}

type Operation struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Step        []string  `json:"step"`
	Time        Time      `json:"time"`
	Frequency   Frequency `json:"frequency"`

	ProgressBar float32 `json:"progress_bar"`
	Achieved    bool    `json:"achieved"`

	// the importance of Operation
	Weight float32 `json:"weight"` // from 0 - 100

	// end recursion if Atom is true
	Atom bool `json:"atom"`
}

// TODO function need three tables
// daily report or weekly/monthly...
// arrange list

type Time struct {
	/**
	can be nil, if Frequency.Oneshot is true
	 */
	ExpectStartTime time.Time     `json:"expect_start_time"`
	ExpectStopTime  time.Time     `json:"expect_stop_time"`
	PredictStopTime time.Time     `json:"predict_stop_time"`
	StartTime       time.Time     `json:"start_time"`
	StopTime        time.Time     `json:"stop_time"`
	ExpectDuration  time.Duration `json:"expect_duration"`
	RealDuration    time.Duration `json:"reality_duration"`
	TimeRate        float32       `json:"time_rate"`
}

type Frequency struct {
	// set true, if just want to do it once
	Oneshot bool `json:"oneshot"`
	// daily, weekly, monthly, quarterly, yearly
	Frequency string `json:"frequency"`
}

func (c *Operation) GetExpectDuration() (expectDuration time.Duration) {
	c.Time.ExpectDuration = c.Time.ExpectStopTime.Sub(c.Time.ExpectStartTime)
	return c.Time.ExpectStopTime.Sub(c.Time.ExpectStartTime)
}

func (c *Operation) GetRealDuration() (realDuration time.Duration) {
	c.Time.RealDuration = c.Time.StopTime.Sub(c.Time.StartTime)
	return c.Time.StopTime.Sub(c.Time.StartTime)
}

func (c *Operation) GetTimeRate() (rate float32) {
	expectdura := float32(c.Time.ExpectDuration.Nanoseconds())
	realdura := float32(c.Time.RealDuration.Nanoseconds())
	rate = expectdura / realdura
	c.Time.TimeRate = rate
	return
}

func (c *Operation) DailyReport() {
	/**
	show things have been done at this day
	 */
	/**
	choose things wanna do at tomorrow
	show progress bar
	 */
	/**
	show things have arranged
	show progress bar
	 */
}
