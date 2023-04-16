package account

import (
    "testing"
)

func TestHashPassword(t *testing.T) {
    const fixturePasswd = "12345"
    const fixtureMod    = "6"

    fstHash := HashPassword(fixturePasswd)
    sndHash := HashPassword(fixturePasswd + fixtureMod)
    thdHash := HashPassword(fixturePasswd)

    if fstHash == sndHash {
        t.Errorf(
            "failed to assert hash inequality\n %s <> %s",
            fstHash, sndHash,
        ) 
    }

    if thdHash != fstHash {
        t.Errorf(
            "failed to assert hash equality\n %s <> %s",
            thdHash, fstHash,
        )
    }

}
