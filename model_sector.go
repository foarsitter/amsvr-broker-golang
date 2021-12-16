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
// Sector struct for Sector
type Sector struct {
	Id int32 `json:"id"`
	Aansluiting int32 `json:"aansluiting"`
	Nummer int32 `json:"nummer"`
	Naam string `json:"naam"`
	Straat string `json:"straat,omitempty"`
	Huisnummer string `json:"huisnummer,omitempty"`
	Huisletter string `json:"huisletter,omitempty"`
	Toevoeging string `json:"toevoeging,omitempty"`
	Postcode string `json:"postcode,omitempty"`
	Plaats string `json:"plaats,omitempty"`
	// https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel
	RijksdriehoekstelX *string `json:"rijksdriehoekstel_x,omitempty"`
	// https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel
	RijksdriehoekstelY *string `json:"rijksdriehoekstel_y,omitempty"`
	Waarschuwingsadressen []Waarschuwingsadres `json:"waarschuwingsadressen"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	// 0 = Test, 1 = Actief
	Toestand *MeldingStatus `json:"toestand,omitempty"`
	Gebruiksfunctie *GebruiksfunctieEnum `json:"gebruiksfunctie,omitempty"`
	BagId string `json:"bag_id"`
	BagLookupId string `json:"bag_lookup_id,omitempty"`
}