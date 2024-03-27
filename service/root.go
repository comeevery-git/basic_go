package service

// Network, repository 의 다리 역할
// (API Reqeust) Network => Service => Repository

import (
	"net/http"

	"example.com/m/network"
	"example.com/m/repository"
	"example.com/m/config"
)
