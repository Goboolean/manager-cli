package command

import (
	"context"
	"errors"
	"regexp"
	"strconv"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// TODO: 함수가 너무 많아지면 옵션 구조체를 넘기는 방안도 생각해보기 클라이언트의 요구사항에 맟춰 변경 가능한 인터페이스
func (a *CommandAdaptor) BackupTrade(ctx context.Context, backupType string, isTransmitted bool) error {
	if backupType == "full" && isTransmitted {
		return a.backUpService.BackupTradeFullToRemote(ctx)
	} else if backupType == "diff" && isTransmitted {
		return a.backUpService.BackupTradeDiffToRemote(ctx)
	} else if backupType == "full" && !isTransmitted {
		return a.backUpService.BackupTradeFull(ctx)
	} else if backupType == "diff" && !isTransmitted {
		return a.backUpService.BackupTradeDiff(ctx)
	}
	return nil
}

func (a *CommandAdaptor) BackupProduct(ctx context.Context, id string, backupType string, isTransmitted bool) error {
	if backupType == "full" && isTransmitted {
		return a.backUpService.BackupProductFullToRemote(ctx, id)
	} else if backupType == "diff" && isTransmitted {
		return a.backUpService.BackupProductDiffToRemote(ctx, id)
	} else if backupType == "full" && !isTransmitted {
		return a.backUpService.BackupProductFull(ctx, id)
	} else if backupType == "diff" && !isTransmitted {
		return a.backUpService.BackupProductDiff(ctx, id)
	}
	return nil
}

type RegisterParms struct {
	Id       string
	Type     string
	Name     string
	Location string
	Exchange string
	Code     string
}

func (a *CommandAdaptor) Register(ctx context.Context, in RegisterParms) error {

	if in.Location == "null" {
		// Blank String indicates null
		// TODO: define blank string as constant
		in.Location = entity.NullString
	}

	return a.regService.RegisterProduct(
		ctx,
		entity.ProductMeta{
			Id:          in.Id,
			Name:        in.Name,
			Code:        in.Code,
			Location:    in.Location,
			Exchange:    in.Exchange,
			Description: entity.NullString,
			Type:        in.Type,
		})
}

// UpdateStatus updates the status information based on the given id and desired value.
// If the desired is number in string, it represents the bit marks of the target status.
// The rightmost bit represents whether a product is relayable(1) or not(0).
// The second bit represents whether a product is being stored or not.
// The third bit represents whether a product is being transmitting or not.
// If the desired  is a string, it is a string literal that includes an operation and status(es).
// Available operators:
//
//	-: remove status(es)
//	+: add status(es)
//	=: set status(es)
//
// Available statuses:
//
//	r: reliable
//	s: stored
//	t: transmitted
func (a *CommandAdaptor) UpdateStatus(ctx context.Context, id string, desired string) error {
	// TODO: Refactor to deduce complexity of control structure

	if matched, _ := regexp.MatchString("^[0-7]{1}$", id); matched {
		// If desired is dial, it is a Bit mark of target status.
		// First bit from the rightmost represents whether a product is relayable(1) or not(0),
		// and Second bit does whether a product is being stored or not,
		// and Third bit whether a product is being transmitting or not

		TargetStatusMask, _ := strconv.ParseInt(desired, 10, 0)

		return a.statusService.SetStatus(
			ctx,
			id,
			entity.ProductStatus{
				Relayable:   TargetStatusMask&1<<2 >= 1,
				Stored:      TargetStatusMask&1<<1 >= 1,
				Transmitted: TargetStatusMask&1<<0 >= 1,
			})

	} else if matched, _ := regexp.MatchString("^(\\+|-|=)(r|s|t){1,3}$", id); matched {

		// If desired is string literal, it consist of operator and status(es)
		// The statues(es) is appended after the operator to indicate the desired action

		arr := []byte(desired)

		TargetStatusEntity := entity.ProductStatus{
			Relayable:   false,
			Stored:      false,
			Transmitted: false,
		}
		// Update the TargetStatusEntity based on the status(es) provided in the desired value.
		for i := 1; i < len(arr); i++ {
			if arr[i] == 'r' {
				TargetStatusEntity.Relayable = true
			} else if arr[i] == 's' {
				TargetStatusEntity.Stored = true
			} else if arr[i] == 't' {
				TargetStatusEntity.Transmitted = true
			}
		}

		if arr[0] == '+' {
			return a.statusService.AddStatus(ctx, id, TargetStatusEntity)
		} else if arr[0] == '-' {
			return a.statusService.RemoveStatus(ctx, id, TargetStatusEntity)
		} else if arr[0] == '=' {
			return a.statusService.SetStatus(ctx, id, TargetStatusEntity)
		}

	} else {
		return errors.New("fail to parse")
	}

	return nil
}

// TODO: Change form of status which api requires
func (a *CommandAdaptor) GetStatus(ctx context.Context, id string) (status string, err error) {
	result, err := a.statusService.GetStatus(ctx, id)

	if err != err {
		status = ""
		return
	}

	if result.Relayable {
		status += "r"
	}

	if result.Stored {
		status += "s"
	}

	if result.Transmitted {
		status += "t"
	}

	return

}
