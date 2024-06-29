package version

import (
	"catalog-service-management-api/internal/domain/models"
	"reflect"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name string
		args struct {
			versions []models.Version
		}
		want []models.Version
	}{
		{
			name: "Simple version sorting",
			args: struct {
				versions []models.Version
			}{
				versions: []models.Version{
					{Number: "1.0.0", CreatedAt: now.Add(-2 * time.Hour)},
					{Number: "1.1.0", CreatedAt: now.Add(-1 * time.Hour)},
					{Number: "0.9.0", CreatedAt: now.Add(-3 * time.Hour)},
				},
			},
			want: []models.Version{
				{Number: "1.1.0", CreatedAt: now.Add(-1 * time.Hour)},
				{Number: "1.0.0", CreatedAt: now.Add(-2 * time.Hour)},
				{Number: "0.9.0", CreatedAt: now.Add(-3 * time.Hour)},
			},
		},
		{
			name: "Complex version sorting",
			args: struct {
				versions []models.Version
			}{
				versions: []models.Version{
					{Number: "2.0.0", CreatedAt: now.Add(-5 * time.Hour)},
					{Number: "1.9.9", CreatedAt: now.Add(-6 * time.Hour)},
					{Number: "2.0.0-alpha", CreatedAt: now.Add(-7 * time.Hour)},
					{Number: "2.0.0-beta", CreatedAt: now.Add(-4 * time.Hour)},
					{Number: "2.0.1", CreatedAt: now.Add(-3 * time.Hour)},
					{Number: "2.1.0-rc.1", CreatedAt: now.Add(-2 * time.Hour)},
					{Number: "2.1.0", CreatedAt: now.Add(-1 * time.Hour)},
				},
			},
			want: []models.Version{
				{Number: "2.1.0", CreatedAt: now.Add(-1 * time.Hour)},
				{Number: "2.1.0-rc.1", CreatedAt: now.Add(-2 * time.Hour)},
				{Number: "2.0.1", CreatedAt: now.Add(-3 * time.Hour)},
				{Number: "2.0.0", CreatedAt: now.Add(-5 * time.Hour)},
				{Number: "2.0.0-beta", CreatedAt: now.Add(-4 * time.Hour)},
				{Number: "2.0.0-alpha", CreatedAt: now.Add(-7 * time.Hour)},
				{Number: "1.9.9", CreatedAt: now.Add(-6 * time.Hour)},
			},
		},
		{
			name: "Single version",
			args: struct {
				versions []models.Version
			}{
				versions: []models.Version{
					{Number: "1.0.0", CreatedAt: now},
				},
			},
			want: []models.Version{
				{Number: "1.0.0", CreatedAt: now},
			},
		},
		{
			name: "Empty versions",
			args: struct {
				versions []models.Version
			}{
				versions: []models.Version{},
			},
			want: []models.Version{},
		},
		{
			name: "Nil versions",
			args: struct {
				versions []models.Version
			}{
				versions: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sort(tt.args.versions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
