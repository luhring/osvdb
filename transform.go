package osvdb

import (
	"github.com/luhring/osvdb/model"
	"github.com/luhring/osvdb/osv"
)

func documentToModel(document osv.Document) model.Vulnerability {
	return model.Vulnerability{
		SchemaVersion:    document.SchemaVersion,
		VulnerabilityID:  document.Id,
		Modified:         document.Modified,
		Published:        document.Published,
		Withdrawn:        document.Withdrawn,
		Aliases:          aliasesToModel(document.Aliases),
		Related:          relatedToModel(document.Related),
		Summary:          document.Summary,
		Details:          document.Details,
		Severity:         severitiesToModel(document.Severity),
		Affected:         affectedToModel(document.Affected),
		References:       referencesToModel(document.References),
		Credits:          creditsToModel(document.Credits),
		DatabaseSpecific: document.DatabaseSpecific,
	}
}

func aliasesToModel(aliases []string) []model.Alias {
	result := make([]model.Alias, 0, len(aliases))
	for _, alias := range aliases {
		result = append(result, model.Alias{Value: alias})
	}
	return result
}

func relatedToModel(related []string) []model.Related {
	result := make([]model.Related, 0, len(related))
	for _, r := range related {
		result = append(result, model.Related{Value: r})
	}
	return result
}

func severitiesToModel(severities []osv.Severity) []model.Severity {
	result := make([]model.Severity, 0, len(severities))
	for _, s := range severities {
		result = append(result, model.Severity{Type: s.Type, Score: s.Score})
	}
	return result
}

func affectedToModel(affected []osv.Affected) []model.Affected {
	result := make([]model.Affected, 0, len(affected))
	for _, a := range affected {
		result = append(result, model.Affected{
			Package:           packageToModel(a.Package),
			Severity:          severitiesToModel(a.Severity),
			Ranges:            rangesToModel(a.Ranges),
			Versions:          versionsToModel(a.Versions),
			EcosystemSpecific: a.EcosystemSpecific,
			DatabaseSpecific:  a.DatabaseSpecific,
		})
	}
	return result
}

func packageToModel(pkg osv.Package) model.Package {
	return model.Package{
		Ecosystem: pkg.Ecosystem,
		Name:      pkg.Name,
		PURL:      pkg.Purl,
	}
}

func rangesToModel(ranges []osv.Range) []model.Range {
	result := make([]model.Range, 0, len(ranges))
	for _, r := range ranges {
		result = append(result, model.Range{
			Type:             r.Type,
			Repo:             r.Repo,
			Events:           eventsToModel(r.Events),
			DatabaseSpecific: r.DatabaseSpecific,
		})
	}
	return result
}

func eventsToModel(events []osv.Event) []model.Event {
	result := make([]model.Event, 0, len(events))
	for _, e := range events {
		result = append(result, model.Event{
			Introduced:   e.Introduced,
			Fixed:        e.Fixed,
			LastAffected: e.LastAffected,
			Limit:        e.Limit,
		})
	}
	return result
}

func versionsToModel(versions []string) []model.Version {
	result := make([]model.Version, 0, len(versions))
	for _, v := range versions {
		result = append(result, model.Version{Value: v})
	}
	return result
}

func referencesToModel(references []osv.Reference) []model.Reference {
	result := make([]model.Reference, 0, len(references))
	for _, r := range references {
		result = append(result, model.Reference{
			Type: r.Type,
			URL:  r.Url,
		})
	}
	return result
}

func creditsToModel(credits []osv.Credit) []model.Credit {
	result := make([]model.Credit, 0, len(credits))
	for _, c := range credits {
		result = append(result, model.Credit{
			Name:    c.Name,
			Contact: contactsToModel(c.Contact),
			Type:    c.Type,
		})
	}
	return result
}

func contactsToModel(contacts []string) []model.Contact {
	result := make([]model.Contact, 0, len(contacts))
	for _, c := range contacts {
		result = append(result, model.Contact{Value: c})
	}
	return result
}
