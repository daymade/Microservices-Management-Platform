package version

import (
	"catalog-service-management-api/internal/domain/models"
	"sort"

	"github.com/Masterminds/semver/v3"
)

// Sort 按语义化版本号降序排序版本，不会修改传入的切片
func Sort(versions []models.Version) []models.Version {
	if versions == nil {
		return nil
	}

	sortedVersions := make([]models.Version, len(versions))
	copy(sortedVersions, versions)

	sort.Slice(sortedVersions, func(i, j int) bool {
		vi, errI := semver.NewVersion(sortedVersions[i].Number)
		vj, errJ := semver.NewVersion(sortedVersions[j].Number)

		if errI != nil {
			return false
		}
		if errJ != nil {
			return true
		}

		return vi.GreaterThan(vj)
	})

	return sortedVersions
}
