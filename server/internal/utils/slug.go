package utils

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
)
var adjectives = []string{
	"swift", "bright", "clever", "mighty", "gentle", "fierce", "noble", "brave",
	"silent", "golden", "silver", "rapid", "bold", "calm", "wild", "free",
	"sharp", "smooth", "strong", "wise", "quick", "dark", "light", "deep",
	"tall", "small", "grand", "pure", "fast", "slow", "loud", "quiet",
}

var nouns = []string{
	"wolf", "eagle", "lion", "tiger", "bear", "hawk", "fox", "owl",
	"shark", "falcon", "panther", "dragon", "phoenix", "raven", "viper",
	"storm", "thunder", "lightning", "river", "mountain", "ocean", "fire",
	"shadow", "blade", "arrow", "stone", "flame", "wind", "star", "moon",
	"sun", "crystal", "diamond", "steel", "iron", "gold", "silver", "bronze",
}

func hash(s string) int {
	h := 0
	for _, char := range s {
		h = ((h << 5) - h) + int(char)
		h &= 0xFFFFFFFF 
	}
	if h < 0 {
		h = -h
	}
	return h
}

func seededRandom(seed int, min int, max int) int {
	x := math.Sin(float64(seed)) * 10000
	frac := x - math.Floor(x)
	return int(frac*float64(max-min+1)) + min
}

func GenerateSlug(ip string, format string, length int, salt ...string) (string, error) {
	fmt.Println("IP",ip)
	if !isValidIP(ip) {
		return "", fmt.Errorf("invalid IP address format")
	}

	baseSeed := hash(ip)
	suffixSeed := baseSeed

	if len(salt) > 0 {
		suffixSeed = hash(ip + salt[0])
	} else {
		suffixSeed = hash(ip + fmt.Sprint(time.Now().UnixNano()))
	}

	switch format {
	case "readable":
		adjIndex := seededRandom(baseSeed, 0, len(adjectives)-1)
		nounIndex := seededRandom(baseSeed*2, 0, len(nouns)-1)
		number := seededRandom(suffixSeed*3, 100, 99999) // larger range to reduce collisions
		return fmt.Sprintf("%s-%s-%d", adjectives[adjIndex], nouns[nounIndex], number), nil

	case "alphanumeric":
		chars := "abcdefghijklmnopqrstuvwxyz0123456789"
		var builder strings.Builder
		for i := 0; i < length; i++ {
			index := seededRandom(suffixSeed+i, 0, len(chars)-1)
			builder.WriteByte(chars[index])
		}
		return builder.String(), nil

	case "hex":
		var builder strings.Builder
		for i := 0; i < length; i++ {
			hexChar := strconv.FormatInt(int64(seededRandom(suffixSeed+i, 0, 15)), 16)
			builder.WriteString(hexChar)
		}
		return builder.String(), nil

	default:
		return "", fmt.Errorf("invalid format: use 'readable', 'alphanumeric', or 'hex'")
	}
}

func isValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}