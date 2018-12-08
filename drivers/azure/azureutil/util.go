package azureutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/* Utilities */

// randomAzureStorageAccountName generates a valid storage account name prefixed
// with a predefined string. Availability of the name is not checked. Uses maximum
// length to maximise randomness.
func randomAzureStorageAccountName() string {
	const (
		maxLen = 24
		chars  = "0123456789abcdefghijklmnopqrstuvwxyz"
	)
	return storageAccountPrefix + randomString(maxLen-len(storageAccountPrefix), chars)
}

// randomString generates a random string of given length using specified alphabet.
func randomString(n int, alphabet string) string {
	r := timeSeed()
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[r.Intn(len(alphabet))]
	}
	return string(b)
}

// imageName holds various components of an OS image name identifier
type imageName struct{ publisher, offer, sku, version string }
 
// parseImageName parses as publisher:offer:sku:version into those parts or /subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Compute/images/<image>
func parseImageName(image string) (imageName, error) {
    if l := strings.Split(image, ":"); len(l) == 4 {
        return imageName{l[0], l[1], l[2], l[3]}, nil
    }
    else if l = strings.Split(image,"/"); len(l) == 8 {
        return imageName{l[0],l[1],l[2],l[3],l[4],l[5],l[6],l[7]}, nil
    }
    else {
        return imageName{}, fmt.Errorf("image name %q not a valid format", image)
    }
}

func timeSeed() *rand.Rand { return rand.New(rand.NewSource(time.Now().UTC().UnixNano())) }
