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
// Aansluiting struct for Aansluiting
type Aansluiting struct {
	Id int32 `json:"id"`
	Naam string `json:"naam"`
	Aansluitnummer string `json:"aansluitnummer"`
	// 0 = Nieuw, 10 = Ingediend ter aanvraag [ATSP], 20 = Aanvraag afgekeurd [Risicobeheer], 30 = Aanvraag goedgekeurd [Risicobeheer], 50 = Live test ingepland [ATSP], 55 = Goedgekeurd [AMS-servicedesk], 60 = Actief gezet [Risicobeheer], 5 = Migratie gestart [ATSP], 61 = Migratie goedgekeurd [AMS-Servicedesk]
	Status AansluitingStatusEnum `json:"status"`
	Atsp int32 `json:"atsp"`
	Gemeente *int32 `json:"gemeente,omitempty"`
	Straat string `json:"straat,omitempty"`
	Huisnummer string `json:"huisnummer,omitempty"`
	Huisletter string `json:"huisletter,omitempty"`
	Toevoeging string `json:"toevoeging,omitempty"`
	Postcode string `json:"postcode,omitempty"`
	Plaats string `json:"plaats,omitempty"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	// 0 = Passief, 1 = Actief
	Toestand MeldingStatus `json:"toestand"`
	BagLookupId string `json:"bag_lookup_id,omitempty"`
	BagId string `json:"bag_id"`
	BagNietBeschikbaar *bool `json:"bag_niet_beschikbaar,omitempty"`
	BagComment string `json:"bag_comment,omitempty"`
	// https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel
	RijksdriehoekstelX *string `json:"rijksdriehoekstel_x,omitempty"`
	// https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel
	RijksdriehoekstelY *string `json:"rijksdriehoekstel_y,omitempty"`
	Gebruiksfunctie *GebruiksfunctieEnum `json:"gebruiksfunctie,omitempty"`
	Bijzonderheden *string `json:"bijzonderheden,omitempty"`
	Bron BronEnum `json:"bron"`
}
