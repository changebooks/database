package database

import "errors"

func NewSchema(p *Profile) (*Schema, error) {
	if p == nil {
		return nil, errors.New("profile can't be nil")
	}

	proto := p.GetProto()
	host := p.GetHost()
	database := p.GetDatabase()
	username := p.GetUsername()
	password := p.GetPassword()
	charset := p.GetCharset()
	collation := p.GetCollation()
	timeout := p.GetTimeout()
	dsn := p.GetDsn()
	port := p.GetPort()
	maxOpen := p.GetMaxOpen()
	maxIdle := p.GetMaxIdle()
	maxLifetime := p.GetMaxLifetime()

	builder := &SchemaBuilder{}
	builder.
		SetProto(proto).
		SetHost(host).
		SetDatabase(database).
		SetUsername(username).
		SetPassword(password).
		SetCharset(charset).
		SetCollation(collation).
		SetTimeout(timeout).
		SetDsn(dsn).
		SetPort(port).
		SetMaxOpen(maxOpen).
		SetMaxIdle(maxIdle).
		SetMaxLifetime(maxLifetime)

	return builder.Build()
}
