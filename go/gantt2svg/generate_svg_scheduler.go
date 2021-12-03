package gantt2svg

import (
	"time"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
)

// GenerateSvgScheduler is the struct of the singloton
// in charge of updating the svg when there is a new commit from the front
type GenerateSvgScheduler struct {
}

var generateSvgSchedulerSingloton GenerateSvgScheduler

// start the scheduler
func init() {
	go generateSvgSchedulerSingloton.checkoutScheduler()
}

func (generateSvgScheduler *GenerateSvgScheduler) checkoutScheduler() {

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// checkoutSchedulerPeriod is the period of the "checkout scheduler"
	var CheckoutSchedulerPeriod = time.NewTicker(100 * time.Millisecond)

	lastPushFromFront := uint(0)
	lastPushFromBack := uint(0)
	for {
		select {
		case t := <-CheckoutSchedulerPeriod.C:

			_ = t

			if gonggantt_models.Stage.BackRepo != nil {
				newPushFromFront := gonggantt_models.Stage.BackRepo.GetLastPushFromFrontNb()
				if lastPushFromFront < newPushFromFront {

					ganttToSVGTranformerSingloton.GenerateSvg(&gonggantt_models.Stage)
					lastPushFromFront = newPushFromFront
				}
			}
			if gonggantt_models.Stage.BackRepo != nil {
				newPushFromBack := gonggantt_models.Stage.BackRepo.GetLastCommitNb()
				if lastPushFromBack < newPushFromBack {

					ganttToSVGTranformerSingloton.GenerateSvg(&gonggantt_models.Stage)
					lastPushFromBack = newPushFromBack
				}
			}
		}
	}
}
