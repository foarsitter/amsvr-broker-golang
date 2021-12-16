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
// ContactPersoon struct for ContactPersoon
type ContactPersoon struct {
	Id int32 `json:"id"`
	Aansluiting int32 `json:"aansluiting"`
	Voornaam string `json:"voornaam"`
	Tussenvoegsel *string `json:"tussenvoegsel,omitempty"`
	Achternaam string `json:"achternaam"`
	IsTechnischContact bool `json:"is_technisch_contact,omitempty"`
	Telefoonnummers []Telefoonnummer `json:"telefoonnummers"`
	// contactpersoon t.b.v. informatie voor de repressieve dienst (verplicht voor een technisch contact)
	Email *string `json:"email,omitempty"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}