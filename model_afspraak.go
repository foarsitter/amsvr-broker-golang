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
// Afspraak struct for Afspraak
type Afspraak struct {
	Id int32 `json:"id"`
	StartTijd time.Time `json:"start_tijd"`
	EindTijd time.Time `json:"eind_tijd,omitempty"`
	Titel string `json:"titel"`
	Title string `json:"title"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	Url string `json:"url"`
	Aansluiting *int32 `json:"aansluiting,omitempty"`
}