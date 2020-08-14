package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"log"
	"test_gqlgen/graph/generated"
	"test_gqlgen/models"
	"time"
)

type Resolver struct{}

func (r *queryResolver) Test(ctx context.Context) (models.TestResult, error) {
	return models.Test{A: 6, B: 7}, nil
}

func (r *queryResolver) Tests(ctx context.Context) (models.TestsResult, error) {
	var tests []*models.Test
	for i := 0; i < 10; i++ {
		tests = append(tests, &models.Test{A: 6 * i, B: 7 * i})
	}

	return models.Tests{Items: tests}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

type testResolver struct{ *Resolver }

type gMetricResolver struct{ *Resolver }

func (r *Resolver) Test() generated.TestResolver {
	return &testResolver{r}
}

func (r *Resolver) GMetric() generated.GMetricResolver {
	return &gMetricResolver{r}
}

func (t *testResolver) G(ctx context.Context, test *models.Test) (*models.G, error) {
	log.Printf(" === G started for T(%d, %d) at %d\n", test.A, test.B, time.Now().UnixNano())
	<-time.After(5 * time.Second)
	log.Printf("G ended for T(%d, %d) at %d ===\n", test.A, test.B, time.Now().UnixNano())
	return &models.G{A: test.A + test.B}, nil
}

func (t *testResolver) H(ctx context.Context, test *models.Test) (*models.G, error) {
	log.Printf(" === H called for T(%d, %d) at %d\n", test.A, test.B, time.Now().UnixNano())
	<-time.After(2 * time.Second)
	log.Printf("H called for T(%d, %d) at %d ===\n", test.A, test.B, time.Now().UnixNano())
	return &models.G{A: test.A + test.B}, nil
}

func (t *testResolver) N(ctx context.Context, test *models.Test) (*models.GMetric, error) {
	log.Println("N resolver called")
	var tests []int
	for i := 0; i < 10; i++ {
		tests = append(tests, 6*i)
	}

	return &models.GMetric{ItemsIds: tests}, nil
}

func (gm *gMetricResolver) Items(ctx context.Context, metric *models.GMetric) ([]*models.G, error) {
	log.Println("Items resolver called")
	var tests []*models.G
	for _, v := range metric.ItemsIds {
		tests = append(tests, &models.G{A: v})
	}

	return tests, nil
}
