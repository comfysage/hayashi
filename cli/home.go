package cli

import "github.com/crispybaccoon/hayashi/pkg"

func Home(name string) error {
  p, err := pkg.GetPkg(name)
  if err != nil {
    return err
  }

  printf(p.Url)

  return nil
}
