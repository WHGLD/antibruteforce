package app

import (
	"context"
	"fmt"
	"net"
)

type Bucket struct {
	login string
	pass  string
	ip    string
}

type App struct {
	limiter    RateLimiter
	repository DBRepository
}

func NewApp(limiter RateLimiter, repository DBRepository) *App {
	app := App{
		limiter:    limiter,
		repository: repository,
	}

	return &app
}

func (a *App) CheckRateLimits(ctx context.Context, login, password, ip string) (bool, error) {
	whiteList, errW := a.repository.GetWhiteList(ctx)
	if errW != nil {
		fmt.Println(errW)
	}
	if checkForWhiteList(ip, whiteList) {
		return true, nil
	}

	blackList, errB := a.repository.GetBlackList(ctx)
	if errB != nil {
		fmt.Println(errB)
	}
	if checkForBlackList(ip, blackList) {
		return false, nil
	}

	return a.limiter.CheckRateLimits(ctx, login, password, ip)
}

func (a *App) Reset(ctx context.Context, login, password, ip string) error {
	return a.limiter.ResetBuckets(ctx, login, password, ip)
}

func (a *App) AddToWhiteList(ctx context.Context, ip, mask string) error {
	// todo проверка ip и mask (подсеть: "192.0.2.0/24")
	return a.repository.AddToWhiteList(ctx, ip, mask)
}

func (a *App) RemoveFromWhiteList(ctx context.Context, ip, mask string) error {
	return a.repository.RemoveFromWhiteList(ctx, ip, mask)
}

func (a *App) AddToBlackList(ctx context.Context, ip, mask string) error {
	// todo проверка ip и mask (подсеть: "192.0.2.0/24")
	return a.repository.AddToBlackList(ctx, ip, mask)
}

func (a *App) RemoveFromBlackList(ctx context.Context, ip, mask string) error {
	return a.repository.RemoveFromBlackList(ctx, ip, mask)
}

func checkForWhiteList(ip string, whiteList []*IPNetDb) bool {
	for _, w := range whiteList {
		IPNet := mapToIPNet(w)
		if IPNet.Contains(net.ParseIP(ip)) {
			return true
		}
	}

	return false
}

func checkForBlackList(ip string, blackList []*IPNetDb) bool {
	for _, w := range blackList {
		IPNet := mapToIPNet(w)
		if IPNet.Contains(net.ParseIP(ip)) {
			return true
		}
	}

	return false
}

func mapToIPNet(IPNetDb *IPNetDb) *net.IPNet {
	byteMask := net.ParseIP(IPNetDb.Mask).To4()
	return &net.IPNet{
		IP:   net.ParseIP(IPNetDb.Ip),
		Mask: net.IPv4Mask(byteMask[0], byteMask[1], byteMask[2], byteMask[3]),
	}
}
