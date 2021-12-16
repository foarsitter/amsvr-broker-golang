/*
 * AMSVR Broker API
 *
 * # Introductie     Voor u heeft u de online documentatie van de AMSVR Broker API. Omdat het een Nederlands domein      betreft worden er Engelse en Nederlandse terminologie door elkaar heen gebruikt.  Liefhebbers van de Swagger UI kunnen [hier terecht](/swagger-ui). Daarnaast is de API eveneens beschikbaar in de op endpoint niveau zoals bijvoorbeeld [hier](/api/aansluitingen).  ##        
 *
 * API version: release-0.6.0.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package broker
import (
	"time"
)
// KiezerVervangenZaakIndienen struct for KiezerVervangenZaakIndienen
type KiezerVervangenZaakIndienen struct {
	Id int32 `json:"id"`
	Aansluiting int32 `json:"aansluiting"`
	// nieuw = Nieuw, ingediend = Ingediend [ATSP], ingepland = Live test ingepland [ATSP], goedgekeurd = Goedgekeurd [AMS-Servicedesk], doorgevoerd = Doorgevoerd [GMS], geannuleerd = Geannuleerd
	Status KiezerVervangenZaakIndienenStatusEnum `json:"status"`
	StatusChanged time.Time `json:"status_changed"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	User int32 `json:"user"`
	Gesloten bool `json:"gesloten"`
	GeslotenDoor int32 `json:"gesloten_door"`
	MogelijkeTransities []string `json:"mogelijke_transities"`
}