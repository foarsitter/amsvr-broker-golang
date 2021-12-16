/*
 * AMSVR Broker API
 *
 * # Introductie     Voor u heeft u de online documentatie van de AMSVR Broker API. Omdat het een Nederlands domein      betreft worden er Engelse en Nederlandse terminologie door elkaar heen gebruikt.  Liefhebbers van de Swagger UI kunnen [hier terecht](/swagger-ui). Daarnaast is de API eveneens beschikbaar in de op endpoint niveau zoals bijvoorbeeld [hier](/api/aansluitingen).  ##        
 *
 * API version: release-0.6.0.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package broker
// AansluitingRequest struct for AansluitingRequest
type AansluitingRequest struct {
	Naam string `json:"naam"`
	Aansluitnummer string `json:"aansluitnummer"`
	Atsp int32 `json:"atsp"`
	Gemeente *int32 `json:"gemeente,omitempty"`
	Straat string `json:"straat,omitempty"`
	Huisnummer string `json:"huisnummer,omitempty"`
	Huisletter string `json:"huisletter,omitempty"`
	Toevoeging string `json:"toevoeging,omitempty"`
	Postcode string `json:"postcode,omitempty"`
	Plaats string `json:"plaats,omitempty"`
	BagLookupId string `json:"bag_lookup_id,omitempty"`
	BagNietBeschikbaar *bool `json:"bag_niet_beschikbaar,omitempty"`
	BagComment string `json:"bag_comment,omitempty"`
	// https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel
	RijksdriehoekstelX *string `json:"rijksdriehoekstel_x,omitempty"`
	// https://www.kadaster.nl/zakelijk/registraties/basisregistraties/rijksdriehoeksmeting/rijksdriehoeksstelsel
	RijksdriehoekstelY *string `json:"rijksdriehoekstel_y,omitempty"`
	Gebruiksfunctie *GebruiksfunctieEnum `json:"gebruiksfunctie,omitempty"`
	Bijzonderheden *string `json:"bijzonderheden,omitempty"`
}
