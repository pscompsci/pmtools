package cache

import "github.com/pscompsci/pmtools/pkg/models"

type engagementCache struct {
	engagements map[string]*models.EngagementSummary
}

func newEngagementCache() *engagementCache {
	return &engagementCache{
		engagements: make(map[string]*models.EngagementSummary),
	}
}

func (c *engagementCache) Insert(engagement *models.EngagementSummary) error {
	entry := c.GetById(engagement.EngagementID)
	if entry != nil {
		return ErrDuplicateEntry
	}
	c.engagements[engagement.EngagementID] = engagement
	return nil
}

func (c *engagementCache) GetById(id string) *models.EngagementSummary {
	if _, exists := c.engagements[id]; !exists {
		return nil
	}
	return c.engagements[id]
}

func (c *engagementCache) GetAll() []*models.EngagementSummary {
	result := []*models.EngagementSummary{}
	for _, summary := range c.engagements {
		result = append(result, summary)
	}
	return result
}

func (c *engagementCache) Update(engagement *models.EngagementSummary) error {
	entry := c.GetById(engagement.EngagementID)
	if entry == nil {
		return ErrNoRecord
	}
	c.engagements[engagement.EngagementID] = engagement
	return nil
}

func (c *engagementCache) Delete(id string) error {
	entry := c.GetById(id)
	if entry == nil {
		return ErrNoRecord
	}
	delete(c.engagements, id)
	return nil
}
