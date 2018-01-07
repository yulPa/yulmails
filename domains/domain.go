package domains

type Environment struct{
  /*
    A environment is a list of domains
  */
  Name string
  domains []string
}

type Allowed struct{
  Env Environment
}

func NewAllowed(env Environment) *Allowed{
  /*
    This function create an Allowed with a given set of environment
    parameter: <environment> A given set of environment
    return: <Allowed> A kind of list of authorized domains
  */
  return &Allowed{
    Env: env,
  }
}
