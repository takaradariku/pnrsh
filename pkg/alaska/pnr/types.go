package pnr

import "encoding/json"

type PNR struct {
	HasGuestPass                    bool            `json:hasGuestPass`
	HasTripInsurance                bool            `json:hasTripInsurance`
	HasNoGoTicketStatus             bool            `json:hasNoGoTicketStatus`
	HasGovernmentFare               bool            `json:hasGovernmentFare`
	IsCompanionFare                 bool            `json:isCompanionFare`
	IsEasyBiz                       bool            `json:isEasyBiz`
	IsDeadHead                      bool            `json:isDeadHead`
	IsAFB                           bool            `json:isAFB`
	IsMinorMissingEmergencyContact  bool            `json:isMinorMissingEmergencyContact`
	IsExpediaVacationPackage        bool            `json:isExpediaVacationPackage`
	IsGroupBooking                  bool            `json:isGroupBooking`
	IsLinkedReservation             bool            `json:isLinkedReservation`
	IsOptedInForFirstClassUpgrade   bool            `json:isOptedInForFirstClassUpgrade`
	IsOptedInForPremiumClassUpgrade bool            `json:isOptedInForPremiumClassUpgrade`
	IsTicketTimeLimitSet            bool            `json:isTicketTimeLimitSet`
	IsUpgradeEligible               bool            `json:isUpgradeEligible`
	IsOnFailureQueue                bool            `json:isOnFailureQueue`
	IsOnPretendItsPurchasedQueue    bool            `json:isOnPretendItsPurchasedQueue`
	IsGoldGuestUpgradeEligible      bool            `json:isGoldGuestUpgradeEligible`
	IsDynamicWaiverEligible         bool            `json:isDynamicWaiverEligible`
	IsDirectorTravel                bool            `json:isDirectorTravel`
	IsIBE                           bool            `json:isIBE`
	ReservationURLs                 json.RawMessage `json:reservationURLs`
	PromoDiscountCode               string          `json:promoDiscountCode`
	GroupBookingID                  json.RawMessage `json:groupBookingId,omitempty`
	TicketTimeLimit                 json.RawMessage `json:ticketTimeLimit,omitempty`
	GroupBookingInfo                json.RawMessage `json:groupBookingInfo,omitempty`
	PaidSsrDetails                  struct {
		HasPetc                    bool `json:hasPetc`
		HasUmnr                    bool `json:hasUmnr`
		HasAvih                    bool `json:hasAvih`
		IsAllPetcPaid              bool `json:isAllPetcPaid`
		IsAllUmnrPaid              bool `json:isAllUmnrPaid`
		IsAllAvihPaid              bool `json:isAllAvihPaid`
		IsUmnrOsiAddedSuccessfully bool `json:isUmnrOsiAddedSuccessfully`
		IsPetcOsiAddedSuccessfully bool `json:isPetcOsiAddedSuccessfully`
		IsAvihOsiAddedSuccessfully bool `json:isAvihOsiAddedSuccessfully`
	} `json:paidSsrDetails`
	Itinerary struct {
		TripType        string `json:tripType`
		Origin          string `json:origin`
		Destination     string `json:destination`
		ItinerarySlices []struct {
			Origin      string `json:origin`
			Destination string `json:destination`
			Segments    []struct {
				IsCheckinEligible              bool            `json:isCheckinEligible`
				IsInSDCWindow                  bool            `json:isInSDCWindow`
				DepartureAirport               string          `json:departureAirport`
				DepartureCity                  string          `json:departureCity`
				DepartureStation               string          `json:departureStation`
				ArrivalAirport                 string          `json:arrivalAirport`
				ArrivalCity                    string          `json:arrivalCity`
				ArrivalStation                 string          `json:arrivalStation`
				MarketingAirlineCode           string          `json:marketingAirlineCode`
				MarketingAirlineName           string          `json:marketingAirlineName`
				MarketingFlightNumber          string          `json:marketingFlightNumber`
				OperatingAirlineCode           string          `json:operatingAirlineCode`
				OperatingAirlineName           string          `json:operatingAirlineName`
				OperatingFlightNumber          string          `json:operatingFlightNumber`
				ActionCode                     string          `json:actionCode`
				DisclosureText                 string          `json:disclosureText`
				CheckinCarrierText             *string         `json:checkinCarrierText`
				ScheduledDepartureDateTime     string          `json:scheduledDepartureDateTime`
				ScheduledDepartureDateTimeUtc  string          `json:scheduledDepartureDateTimeUtc`
				EstimatedDepartureDateTime     *string         `json:estimatedDepartureDateTime`
				EstimatedDepartureDateTimeUtc  *string         `json:estimatedDepartureDateTimeUtc`
				ScheduledArrivalDateTime       string          `json:scheduledArrivalDateTime`
				ScheduledArrivalDateTimeUtc    string          `json:scheduledArrivalDateTimeUtc`
				EstimatedArrivalDateTime       *string         `json:estimatedArrivalDateTime`
				EstimatedArrivalDateTimeUtc    *string         `json:estimatedArrivalDateTimeUtc`
				ScheduledDurationInMinutes     int             `json:scheduledDurationInMinutes`
				EstimatedDurationInMinutes     *int            `json:estimatedDurationInMinutes`
				Status                         *string         `json:status`
				FlightStatus                   *string         `json:flightStatus`
				Cabin                          string          `json:cabin`
				EquipmentName                  string          `json:equipmentName`
				Distance                       int             `json:distance`
				ClassOfService                 *string         `json:classOfService`
				IsArnk                         bool            `json:isArnk`
				SpecialServiceRequests         []APISSR        `json:specialServiceRequests`
				HiddenStops                    json.RawMessage `json:hiddenStops`
				Sequence                       int             `json:sequence`
				IsDisrupted                    bool            `json:isDisrupted`
				IsFlown                        bool            `json:isFlown`
				IsWaitingUpgradeFirstClass     bool            `json:isWaitingUpgradeFirstClass`
				IsWaitingUpgradePremiumClass   bool            `json:isWaitingUpgradePremiumClass`
				IsConfirmedUpgradeFirstClass   bool            `json:isConfirmedUpgradeFirstClass`
				IsConfirmedUpgradePremiumClass bool            `json:isConfirmedUpgradePremiumClass`
				IsStandBy                      bool            `json:isStandBy`
				IsNonRevConfirmed              bool            `json:isNonRevConfirmed`
				IsPetcPaid                     bool            `json:isPetcPaid`
				IsUmnrPaid                     bool            `json:isUmnrPaid`
				IsAvihPaid                     bool            `json:isAvihPaid`
				IsInternationalFlight          bool            `json:isInternationalFlight`
				OtherAirlineRecordLocator      *string         `json:otherAirlineRecordLocator`
				IsGoldGuestUpgradeEligible     bool            `json:isGoldGuestUpgradeEligible`
				GguRuleViolations              *struct {
					IsSegmentNotAAG                      bool `json:isSegmentNotAAG`
					IsArnk                               bool `json:isArnk`
					IsClassOfServiceIneligibleForUpgrade bool `json:isClassOfServiceIneligibleForUpgrade`
					IsAlreadyFirstClass                  bool `json:isAlreadyFirstClass`
					IsPastFlight                         bool `json:isPastFlight`
				} `json:gguRuleViolations`

				Brand string `json:brand`
			} `json:segments`
			PreviousItinerary          json.RawMessage `json:previousItinerary`
			HistoricItineraries        json.RawMessage `json:historicItineraries`
			IsInternationalTrip        bool            `json:isInternationalTrip`
			IsBagCheckedIn             bool            `json:isBagCheckedIn`
			IsBagTagCreated            bool            `json:isBagTagCreated`
			IsOAInitiatedTrip          bool            `json:isOAInitiatedTrip`
			IsInSDCWindow              bool            `json:isInSDCWindow`
			HasScheduleChange          bool            `json:hasScheduleChange`
			DisruptionType             string          `json:disruptionType`
			IsGoldGuestUpgradeEligible bool            `json:isGoldGuestUpgradeEligible`
		} `json:itinerarySlices`
		MatchesTickets         bool   `json:matchesTickets`
		IsRefunded             bool   `json:isRefunded`
		OriginCountryCode      string `json:originCountryCode`
		DestinationCountryCode string `json:destinationCountryCode`
	} `json:itinerary`
	Passengers []struct {
		FirstName               string      `json:firstName`
		MiddleInitial           *string     `json:middleInitial`
		LastName                string      `json:lastName`
		NameRefNumber           string      `json:nameRefNumber`
		AgeGroup                string      `json:ageGroup`
		TierStatus              *string     `json:tierStatus`
		LoyaltyNumber           *string     `json:loyaltyNumber`
		IsCheckedIn             bool        `json:isCheckedIn`
		Tickets                 []APITicket `json:tickets`
		SeatNumbers             [][]*string `json:seatNumbers`
		SpecialServiceRequests  []APISSR    `json:specialServiceRequests`
		OSIs                    []APIOSI    `json:osis`
		TicketsMatchesItinerary bool        `json:ticketsMatchesItinerary`
		KnownTravelerNumber     *string     `json:knownTravelerNumber`
		HasEmergencyContact     bool        `json:hasEmergencyContact`
		HasSecureFlightInfo     bool        `json:hasSecureFlightInfo`
		IsTravelDocEntered      bool        `json:isTravelDocEntered`
		RedressNumber           *string     `json:redressNumber`
		Loyalty                 *struct {
			AirlineCode        string `json:airlineCode`
			AirlineName        string `json:airlineName`
			LoyaltyNumber      string `json:loyaltyNumber`
			AlaskaTierStatus   string `json:alaskaTierStatus`
			OneWorldTierStatus string `json:oneWorldTierStatus`
		} `json:loyalty`
		IsLapInfant      bool `json:isLapInfant`
		ActionsAvailable struct {
			FrequentFlyerNumber      bool `json:frequentFlyerNumber`
			KnownTravelerNumber      bool `json:knownTravelerNumber`
			RedressNumber            bool `json:redressNumber`
			AddSpecialServiceRequest bool `json:addSpecialServiceRequest`
		} `json:actionsAvailable`
		SeatedWith json.RawMessage `json:seatedWith`
	} `json:passengers`
	OSIs                             []APIOSI        `json:osis`
	Tickets                          []APITicket     `json:tickets`
	LinkedReservations               json.RawMessage `json:linkedReservations`
	ReservationScheduleChangeDetails struct {
		HasScheduleChange              bool `json:hasScheduleChange`
		HasPreviousMajorScheduleChange bool `json:hasPreviousMajorScheduleChange`
		FlightChangeQualified          bool `json:flightChangeQualified`
	} `json:reservationScheduleChangeDetails`
	DapUpdateTimestamp string `json:dapUpdateTimestamp`
}

type APISSR struct {
	ID                 string  `json:id`
	ServiceCode        string  `json:serviceCode`
	FreeText           string  `json:freeText`
	ServiceDescription *string `json:serviceDescription`
	ActionCode         string  `json:actionCode`
	VendorCode         string  `json:vendorCode`
	FlightNumber       *string `json:flightNumber`
	FlightDate         string  `json:flightDate`
	Origin             *string `json:origin`
	Destination        *string `json:destination`
}

type APIOSI struct {
	ID         string `json:id`
	FreeText   string `json:freeText`
	FullText   string `json:fullText`
	VendorCode string `json:vendorCode`
}

type APITicket struct {
	Number     string  `json:number`
	Type       string  `json:type`
	Designator *string `json:designator`
	Coupons    []struct {
		Index                 int    `json:index`
		DepartureDateTime     string `json:departureDateTime`
		Origin                string `json:origin`
		Destination           string `json:destination`
		FareBasis             string `json:fareBasis`
		Status                string `json:status`
		MarketingFlightNumber string `json:marketingFlightNumber`
	} `json:coupons`
	Payments []struct {
		Type        string `json:type`
		Certificate string `json:certificate`
	} `json:payments`
	PassengerNameNumber  *string `json:passengerNameNumber`
	FirstName            string  `json:firstName`
	LastName             string  `json:lastName`
	IsActive             bool    `json:isActive`
	ConfirmationCode     string  `json:confirmationCode`
	IsPurged             bool    `json:isPurged`
	SystemCreateDateTime string  `json:systemCreateDateTime`
	HasNonRefEndorsement bool    `json:hasNonRefEndorsement`
}
