package shop

// VAT is the French VAT rate, in percent.
const VAT = 20

// WithVAT returns p with the VAT added.
func WithVAT(p Price) Price { return p + p*VAT/100 }
