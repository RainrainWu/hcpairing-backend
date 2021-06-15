package hcpairing

const (

	// Specialties
	Anaesthesiology          = "Anaesthesiology"
	AllergologyImmunology    = "AllergologyImmunology"
	Biochemistry             = "Biochemistry"
	Cardiology               = "Cardiology"
	Cardiovascular           = "Cardiovascular/Vascular Surgery"
	ChildDentistry           = "Child Dentistry (pedodonty)"
	ClinicalResearch         = "Clinical Research"
	Dermatology              = "Dermatology"
	Dentistry                = "Dentistry"
	DentalHygiene            = "Dental Hygiene"
	Diabetology              = "Diabetology"
	EmergencyMedicine        = "Emergency Medicine"
	Endocrinology            = "Endocrinology"
	FamilyMedicine           = "Family Medicine"
	Gastroenterology         = "Gastroenterology"
	GeneralMedicine          = "General Medicine"
	GeneralSurgery           = "General Surgery"
	Geriatrics               = "Geriatrics"
	MedicalInternship        = "Medical Internship"
	Microbiology             = "Microbiology"
	MagneticResonanceImaging = "Magnetic Resonance Imaging"
	Nephrology               = "Nephrology"
	Neurology                = "Neurology"
	Neurosurgery             = "Neurosurgery"
	Neuropathology           = "Neuropathology"
	NuclearMedicine          = "Nuclear Medicine"
	Obstetrics               = "Obstetrics/Gynecology"
	OccupationalHealth       = "Occupational Health"
	OccupationalTherapy      = "Occupational Therapy"
	Ophthalmology            = "Ophthalmology"
	Oncology                 = "Oncology"
	OrthopaedicSurgery       = "Orthopaedic Surgery"
	Orthodontics             = "Orthodontics"
	Otorhinolaryngology      = "Otorhinolaryngology"
	Paediatrics              = "Paediatrics"
	PhysicalTherapy          = "Physical Therapy"
	PlasticSurgery           = "Plastic Surgery"
	Psychiatry               = "Psychiatry"
	PublicHealth             = "Public Health"
	Pneumology               = "Pneumology"
	Radiology                = "Radiology"
	Research                 = "Research"
	RespiratoryTherapy       = "Respiratory Therapy"
	Rheumatology             = "Rheumatology"
	Urology                  = "Urology"
)

var (
	directMapping map[string][]string = map[string][]string{
		Anxiety:           {Neurology},
		Bloated:           {Gastroenterology},
		BlurredVision:     {Ophthalmology},
		Burn:              {Dermatology},
		ChestFeelsTight:   {Cardiology},
		Constipation:      {Gastroenterology},
		Cough:             {Pneumology, FamilyMedicine, Cardiology, Otorhinolaryngology},
		Diarrhea:          {Neurology, Gastroenterology},
		Drowsy:            {Neurology},
		DrummingInTheEars: {Otorhinolaryngology},
		Fatigue:           {Neurology},
		Fever:             {FamilyMedicine},
		Fracture:          {OrthopaedicSurgery, GeneralSurgery},
		Headache:          {FamilyMedicine, Neurology, Ophthalmology},
		Heatstroke:        {Nephrology},
		Insomnia:          {Neurology},
		ItchySkin:         {Dermatology},
		LossOfHearing:     {Otorhinolaryngology},
		LostAppetite:      {Neurology, Gastroenterology, FamilyMedicine},
		NasalBleeding:     {Otorhinolaryngology, FamilyMedicine},
		NasalDischarge:    {Otorhinolaryngology, FamilyMedicine},
		Pregnancy:         {Obstetrics},
		Rash:              {Dermatology, AllergologyImmunology, FamilyMedicine},
		RunnyNose:         {Otorhinolaryngology, FamilyMedicine},
		ShortOfBreath:     {Cardiology, FamilyMedicine, Otorhinolaryngology},
		Sneeze:            {Otorhinolaryngology, FamilyMedicine},
		SoreEyes:          {Ophthalmology},
		SoreMuscles:       {FamilyMedicine},
		SoreThroat:        {Otorhinolaryngology, FamilyMedicine, Dentistry, ChildDentistry},
		Stomachache:       {Gastroenterology, GeneralSurgery},
		Stuffy:            {FamilyMedicine, Otorhinolaryngology},
		Toothache:         {Dentistry, ChildDentistry},
		Vomit:             {Neurology, Gastroenterology, FamilyMedicine},
	}
)

func contains(arr []string, str string) bool {

	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func setsLeftMerge(set1, set2 []string) []string {

	for _, item := range set2 {
		if !contains(set1, item) {
			set1 = append(set1, item)
		}
	}
	return set1
}

func DirectConversion(tags []string, limit int) []string {

	specialties := []string{}
	for _, tag := range tags {
		specialties = setsLeftMerge(specialties, directMapping[tag])
	}
	return specialties
}
