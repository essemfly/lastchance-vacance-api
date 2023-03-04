package crawler

import (
	"testing"

	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/internal/domain"
)

func TestCrawlPage(t *testing.T) {
	cmd.InitBase()

	hideIdx := 530520001
	soldoutIdx := 519218848
	saleIdx := 362459007
	removedIdx := 362459009

	_, err := CrawlPage(hideIdx)
	if err == nil {
		t.Errorf("hide product should emit error")
	}

	soldoutPd, err := CrawlPage(soldoutIdx)
	if err != nil {
		t.Errorf("soldout product should not emit error")
	}
	if soldoutPd.Status != domain.DANGGN_STATUS_SOLDOUT {
		t.Errorf("soldout product should be soldout status")
	}

	salePd, err := CrawlPage(saleIdx)
	if err != nil {
		t.Errorf("sale product should not emit error")
	}
	if salePd.Status != domain.DANGGN_STATUS_SALE {
		t.Errorf("sale product should be sale status")
	}

	_, err = CrawlPage(removedIdx)
	if err == nil {
		t.Errorf("removed product should emit error")
	}
}
