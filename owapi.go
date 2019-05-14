package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OWPlayer struct {
	CompetitiveStats struct {
		Awards struct {
			Cards        int `json:"cards"`
			Medals       int `json:"medals"`
			MedalsBronze int `json:"medalsBronze"`
			MedalsSilver int `json:"medalsSilver"`
			MedalsGold   int `json:"medalsGold"`
		} `json:"awards"`
		Games struct {
			Played int `json:"played"`
			Won    int `json:"won"`
		} `json:"games"`
	} `json:"competitiveStats"`
	Endorsement     int    `json:"endorsement"`
	EndorsementIcon string `json:"endorsementIcon"`
	GamesWon        int    `json:"gamesWon"`
	Icon            string `json:"icon"`
	Level           int    `json:"level"`
	LevelIcon       string `json:"levelIcon"`
	Name            string `json:"name"`
	Prestige        int    `json:"prestige"`
	PrestigeIcon    string `json:"prestigeIcon"`
	Private         bool   `json:"private"`
	QuickPlayStats  struct {
		Awards struct {
			Cards        int `json:"cards"`
			Medals       int `json:"medals"`
			MedalsBronze int `json:"medalsBronze"`
			MedalsSilver int `json:"medalsSilver"`
			MedalsGold   int `json:"medalsGold"`
		} `json:"awards"`
		Games struct {
			Played int `json:"played"`
			Won    int `json:"won"`
		} `json:"games"`
	} `json:"quickPlayStats"`
	Rating     int    `json:"rating"`
	RatingIcon string `json:"ratingIcon"`
}

type OWCompleteStats struct {
	CompetitiveStats struct {
		Awards struct {
			Cards        int `json:"cards"`
			Medals       int `json:"medals"`
			MedalsBronze int `json:"medalsBronze"`
			MedalsSilver int `json:"medalsSilver"`
			MedalsGold   int `json:"medalsGold"`
		} `json:"awards"`
		CareerStats struct {
			AllHeroes struct {
				Assists struct {
					DefensiveAssists int `json:"defensiveAssists"`
					HealingDone      int `json:"healingDone"`
					OffensiveAssists int `json:"offensiveAssists"`
				} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
					BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
					DeathsAvgPer10Min            float64 `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min      float64 `json:"eliminationsAvgPer10Min"`
					FinalBlowsAvgPer10Min        float64 `json:"finalBlowsAvgPer10Min"`
					HealingDoneAvgPer10Min       float64 `json:"healingDoneAvgPer10Min"`
					HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
					ObjectiveKillsAvgPer10Min    float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min     string  `json:"objectiveTimeAvgPer10Min"`
					SoloKillsAvgPer10Min         int     `json:"soloKillsAvgPer10Min"`
					TimeSpentOnFireAvgPer10Min   string  `json:"timeSpentOnFireAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
					BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
					DefensiveAssistsMostInGame  int    `json:"defensiveAssistsMostInGame"`
					EliminationsMostInGame      int    `json:"eliminationsMostInGame"`
					FinalBlowsMostInGame        int    `json:"finalBlowsMostInGame"`
					HealingDoneMostInGame       int    `json:"healingDoneMostInGame"`
					HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
					KillsStreakBest             int    `json:"killsStreakBest"`
					MeleeFinalBlowsMostInGame   int    `json:"meleeFinalBlowsMostInGame"`
					ObjectiveKillsMostInGame    int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame     string `json:"objectiveTimeMostInGame"`
					OffensiveAssistsMostInGame  int    `json:"offensiveAssistsMostInGame"`
					SoloKillsMostInGame         int    `json:"soloKillsMostInGame"`
					TimeSpentOnFireMostInGame   string `json:"timeSpentOnFireMostInGame"`
					TurretsDestroyedMostInGame  int    `json:"turretsDestroyedMostInGame"`
				} `json:"best"`
				Combat struct {
					BarrierDamageDone int    `json:"barrierDamageDone"`
					DamageDone        int    `json:"damageDone"`
					Deaths            int    `json:"deaths"`
					Eliminations      int    `json:"eliminations"`
					FinalBlows        int    `json:"finalBlows"`
					HeroDamageDone    int    `json:"heroDamageDone"`
					MeleeFinalBlows   int    `json:"meleeFinalBlows"`
					ObjectiveKills    int    `json:"objectiveKills"`
					ObjectiveTime     string `json:"objectiveTime"`
					SoloKills         int    `json:"soloKills"`
					TimeSpentOnFire   string `json:"timeSpentOnFire"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific interface{} `json:"heroSpecific"`
				Game         struct {
					GamesLost   int    `json:"gamesLost"`
					GamesPlayed int    `json:"gamesPlayed"`
					GamesWon    int    `json:"gamesWon"`
					TimePlayed  string `json:"timePlayed"`
				} `json:"game"`
				MatchAwards struct {
					Cards        int `json:"cards"`
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsGold   int `json:"medalsGold"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous struct {
					TurretsDestroyed int `json:"turretsDestroyed"`
				} `json:"miscellaneous"`
			} `json:"allHeroes"`
			Ana struct {
				Assists struct {
					DefensiveAssists            int     `json:"defensiveAssists"`
					DefensiveAssistsAvgPer10Min float64 `json:"defensiveAssistsAvgPer10Min"`
					DefensiveAssistsMostInGame  int     `json:"defensiveAssistsMostInGame"`
					HealingDone                 int     `json:"healingDone"`
					HealingDoneAvgPer10Min      float64 `json:"healingDoneAvgPer10Min"`
					HealingDoneMostInGame       int     `json:"healingDoneMostInGame"`
					OffensiveAssists            int     `json:"offensiveAssists"`
					OffensiveAssistsAvgPer10Min float64 `json:"offensiveAssistsAvgPer10Min"`
					OffensiveAssistsMostInGame  int     `json:"offensiveAssistsMostInGame"`
				} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
					BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
					DeathsAvgPer10Min            float64 `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min      float64 `json:"eliminationsAvgPer10Min"`
					EliminationsPerLife          float64 `json:"eliminationsPerLife"`
					FinalBlowsAvgPer10Min        int     `json:"finalBlowsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
					ObjectiveKillsAvgPer10Min    float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min     string  `json:"objectiveTimeAvgPer10Min"`
					SoloKillsAvgPer10Min         int     `json:"soloKillsAvgPer10Min"`
					TimeSpentOnFireAvgPer10Min   string  `json:"timeSpentOnFireAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife     int    `json:"allDamageDoneMostInLife"`
					BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
					EliminationsMostInGame      int    `json:"eliminationsMostInGame"`
					EliminationsMostInLife      int    `json:"eliminationsMostInLife"`
					FinalBlowsMostInGame        int    `json:"finalBlowsMostInGame"`
					HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife    int    `json:"heroDamageDoneMostInLife"`
					KillsStreakBest             int    `json:"killsStreakBest"`
					ObjectiveKillsMostInGame    int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame     string `json:"objectiveTimeMostInGame"`
					SoloKillsMostInGame         int    `json:"soloKillsMostInGame"`
					TimeSpentOnFireMostInGame   string `json:"timeSpentOnFireMostInGame"`
					WeaponAccuracyBestInGame    string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					BarrierDamageDone  int    `json:"barrierDamageDone"`
					DamageDone         int    `json:"damageDone"`
					Deaths             int    `json:"deaths"`
					Eliminations       int    `json:"eliminations"`
					FinalBlows         int    `json:"finalBlows"`
					HeroDamageDone     int    `json:"heroDamageDone"`
					ObjectiveKills     int    `json:"objectiveKills"`
					ObjectiveTime      string `json:"objectiveTime"`
					QuickMeleeAccuracy string `json:"quickMeleeAccuracy"`
					SoloKills          int    `json:"soloKills"`
					TimeSpentOnFire    string `json:"timeSpentOnFire"`
					WeaponAccuracy     string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					BioticGrenadeKills           int     `json:"bioticGrenadeKills"`
					EnemiesSlept                 int     `json:"enemiesSlept"`
					EnemiesSleptAvgPer10Min      float64 `json:"enemiesSleptAvgPer10Min"`
					EnemiesSleptMostInGame       int     `json:"enemiesSleptMostInGame"`
					HealingAmplified             int     `json:"healingAmplified"`
					HealingAmplifiedAvgPer10Min  float64 `json:"healingAmplifiedAvgPer10Min"`
					HealingAmplifiedMostInGame   int     `json:"healingAmplifiedMostInGame"`
					NanoBoostAssists             int     `json:"nanoBoostAssists"`
					NanoBoostAssistsAvgPer10Min  int     `json:"nanoBoostAssistsAvgPer10Min"`
					NanoBoostAssistsMostInGame   int     `json:"nanoBoostAssistsMostInGame"`
					NanoBoostsApplied            int     `json:"nanoBoostsApplied"`
					NanoBoostsAppliedAvgPer10Min float64 `json:"nanoBoostsAppliedAvgPer10Min"`
					NanoBoostsAppliedMostInGame  int     `json:"nanoBoostsAppliedMostInGame"`
					ScopedAccuracy               string  `json:"scopedAccuracy"`
					ScopedAccuracyBestInGame     string  `json:"scopedAccuracyBestInGame"`
					SecondaryFireAccuracy        string  `json:"secondaryFireAccuracy"`
					SelfHealing                  int     `json:"selfHealing"`
					SelfHealingAvgPer10Min       float64 `json:"selfHealingAvgPer10Min"`
					SelfHealingMostInGame        int     `json:"selfHealingMostInGame"`
					UnscopedAccuracy             string  `json:"unscopedAccuracy"`
					UnscopedAccuracyBestInGame   string  `json:"unscopedAccuracyBestInGame"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost     int    `json:"gamesLost"`
					GamesPlayed   int    `json:"gamesPlayed"`
					GamesWon      int    `json:"gamesWon"`
					TimePlayed    string `json:"timePlayed"`
					WinPercentage string `json:"winPercentage"`
				} `json:"game"`
				MatchAwards struct {
					Cards        int `json:"cards"`
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsGold   int `json:"medalsGold"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous struct {
					TurretsDestroyed int `json:"turretsDestroyed"`
				} `json:"miscellaneous"`
			} `json:"ana"`
			Baptiste struct {
				Assists struct {
					DefensiveAssists            int     `json:"defensiveAssists"`
					DefensiveAssistsAvgPer10Min float64 `json:"defensiveAssistsAvgPer10Min"`
					DefensiveAssistsMostInGame  int     `json:"defensiveAssistsMostInGame"`
					HealingDone                 int     `json:"healingDone"`
					HealingDoneAvgPer10Min      float64 `json:"healingDoneAvgPer10Min"`
					HealingDoneMostInGame       int     `json:"healingDoneMostInGame"`
					OffensiveAssists            int     `json:"offensiveAssists"`
					OffensiveAssistsAvgPer10Min float64 `json:"offensiveAssistsAvgPer10Min"`
					OffensiveAssistsMostInGame  int     `json:"offensiveAssistsMostInGame"`
				} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
					BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
					CriticalHitsAvgPer10Min      float64 `json:"criticalHitsAvgPer10Min"`
					DeathsAvgPer10Min            int     `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min      float64 `json:"eliminationsAvgPer10Min"`
					EliminationsPerLife          int     `json:"eliminationsPerLife"`
					FinalBlowsAvgPer10Min        int     `json:"finalBlowsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
					MeleeFinalBlowsAvgPer10Min   int     `json:"meleeFinalBlowsAvgPer10Min"`
					ObjectiveKillsAvgPer10Min    float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min     string  `json:"objectiveTimeAvgPer10Min"`
					TimeSpentOnFireAvgPer10Min   string  `json:"timeSpentOnFireAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife     int    `json:"allDamageDoneMostInLife"`
					BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
					CriticalHitsMostInGame      int    `json:"criticalHitsMostInGame"`
					CriticalHitsMostInLife      int    `json:"criticalHitsMostInLife"`
					EliminationsMostInGame      int    `json:"eliminationsMostInGame"`
					EliminationsMostInLife      int    `json:"eliminationsMostInLife"`
					FinalBlowsMostInGame        int    `json:"finalBlowsMostInGame"`
					HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife    int    `json:"heroDamageDoneMostInLife"`
					KillsStreakBest             int    `json:"killsStreakBest"`
					MeleeFinalBlowsMostInGame   int    `json:"meleeFinalBlowsMostInGame"`
					ObjectiveKillsMostInGame    int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame     string `json:"objectiveTimeMostInGame"`
					TimeSpentOnFireMostInGame   string `json:"timeSpentOnFireMostInGame"`
					WeaponAccuracyBestInGame    string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					BarrierDamageDone    int    `json:"barrierDamageDone"`
					CriticalHits         int    `json:"criticalHits"`
					CriticalHitsAccuracy string `json:"criticalHitsAccuracy"`
					DamageDone           int    `json:"damageDone"`
					Deaths               int    `json:"deaths"`
					Eliminations         int    `json:"eliminations"`
					FinalBlows           int    `json:"finalBlows"`
					HeroDamageDone       int    `json:"heroDamageDone"`
					MeleeFinalBlows      int    `json:"meleeFinalBlows"`
					ObjectiveKills       int    `json:"objectiveKills"`
					ObjectiveTime        string `json:"objectiveTime"`
					QuickMeleeAccuracy   string `json:"quickMeleeAccuracy"`
					TimeSpentOnFire      string `json:"timeSpentOnFire"`
					WeaponAccuracy       string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					AmplificationMatrixAssists                 int     `json:"amplificationMatrixAssists"`
					AmplificationMatrixAssistsAvgPer10Min      float64 `json:"amplificationMatrixAssistsAvgPer10Min"`
					AmplificationMatrixAssistsBestInGame       int     `json:"amplificationMatrixAssistsBestInGame"`
					AmplificationMatrixCasts                   int     `json:"amplificationMatrixCasts"`
					AmplificationMatrixCastsAvgPer10Min        float64 `json:"amplificationMatrixCastsAvgPer10Min"`
					AmplificationMatrixCastsMostInGame         int     `json:"amplificationMatrixCastsMostInGame"`
					DamageAmplified                            int     `json:"damageAmplified"`
					DamageAmplifiedAvgPer10Min                 float64 `json:"damageAmplifiedAvgPer10Min"`
					DamageAmplifiedMostInGame                  int     `json:"damageAmplifiedMostInGame"`
					HealingAccuracy                            string  `json:"healingAccuracy"`
					HealingAccuracyBestInGame                  string  `json:"healingAccuracyBestInGame"`
					HealingAmplified                           int     `json:"healingAmplified"`
					HealingAmplifiedAvgPer10Min                float64 `json:"healingAmplifiedAvgPer10Min"`
					HealingAmplifiedMostInGame                 int     `json:"healingAmplifiedMostInGame"`
					ImmortalityFieldDeathsPrevented            int     `json:"immortalityFieldDeathsPrevented"`
					ImmortalityFieldDeathsPreventedAvgPer10Min float64 `json:"immortalityFieldDeathsPreventedAvgPer10Min"`
					ImmortalityFieldDeathsPreventedMostInGame  int     `json:"immortalityFieldDeathsPreventedMostInGame"`
					SecondaryFireAccuracy                      string  `json:"secondaryFireAccuracy"`
					SelfHealing                                int     `json:"selfHealing"`
					SelfHealingAvgPer10Min                     float64 `json:"selfHealingAvgPer10Min"`
					SelfHealingMostInGame                      int     `json:"selfHealingMostInGame"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost     int    `json:"gamesLost"`
					GamesPlayed   int    `json:"gamesPlayed"`
					GamesWon      int    `json:"gamesWon"`
					TimePlayed    string `json:"timePlayed"`
					WinPercentage string `json:"winPercentage"`
				} `json:"game"`
				MatchAwards struct {
					Cards        int `json:"cards"`
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsGold   int `json:"medalsGold"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous interface{} `json:"miscellaneous"`
			} `json:"baptiste"`
			DVa struct {
				Assists interface{} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
					BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
					CriticalHitsAvgPer10Min      float64 `json:"criticalHitsAvgPer10Min"`
					DeathsAvgPer10Min            int     `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min      float64 `json:"eliminationsAvgPer10Min"`
					EliminationsPerLife          float64 `json:"eliminationsPerLife"`
					FinalBlowsAvgPer10Min        float64 `json:"finalBlowsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
					MeleeFinalBlowsAvgPer10Min   int     `json:"meleeFinalBlowsAvgPer10Min"`
					ObjectiveKillsAvgPer10Min    float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min     string  `json:"objectiveTimeAvgPer10Min"`
					SoloKillsAvgPer10Min         int     `json:"soloKillsAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife     int    `json:"allDamageDoneMostInLife"`
					BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
					CriticalHitsMostInGame      int    `json:"criticalHitsMostInGame"`
					CriticalHitsMostInLife      int    `json:"criticalHitsMostInLife"`
					EliminationsMostInGame      int    `json:"eliminationsMostInGame"`
					EliminationsMostInLife      int    `json:"eliminationsMostInLife"`
					FinalBlowsMostInGame        int    `json:"finalBlowsMostInGame"`
					HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife    int    `json:"heroDamageDoneMostInLife"`
					KillsStreakBest             int    `json:"killsStreakBest"`
					MeleeFinalBlowsMostInGame   int    `json:"meleeFinalBlowsMostInGame"`
					ObjectiveKillsMostInGame    int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame     string `json:"objectiveTimeMostInGame"`
					SoloKillsMostInGame         int    `json:"soloKillsMostInGame"`
					WeaponAccuracyBestInGame    string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					BarrierDamageDone    int    `json:"barrierDamageDone"`
					CriticalHits         int    `json:"criticalHits"`
					CriticalHitsAccuracy string `json:"criticalHitsAccuracy"`
					DamageDone           int    `json:"damageDone"`
					Deaths               int    `json:"deaths"`
					Eliminations         int    `json:"eliminations"`
					FinalBlows           int    `json:"finalBlows"`
					HeroDamageDone       int    `json:"heroDamageDone"`
					MeleeFinalBlows      int    `json:"meleeFinalBlows"`
					ObjectiveKills       int    `json:"objectiveKills"`
					ObjectiveTime        string `json:"objectiveTime"`
					QuickMeleeAccuracy   string `json:"quickMeleeAccuracy"`
					SoloKills            int    `json:"soloKills"`
					WeaponAccuracy       string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					DamageBlocked                int     `json:"damageBlocked"`
					DamageBlockedAvgPer10Min     float64 `json:"damageBlockedAvgPer10Min"`
					DamageBlockedMostInGame      int     `json:"damageBlockedMostInGame"`
					MechDeaths                   int     `json:"mechDeaths"`
					MechsCalled                  int     `json:"mechsCalled"`
					MechsCalledAvgPer10Min       float64 `json:"mechsCalledAvgPer10Min"`
					MechsCalledMostInGame        int     `json:"mechsCalledMostInGame"`
					SecondaryFireAccuracy        string  `json:"secondaryFireAccuracy"`
					SelfDestructKills            int     `json:"selfDestructKills"`
					SelfDestructKillsAvgPer10Min int     `json:"selfDestructKillsAvgPer10Min"`
					SelfDestructKillsMostInGame  int     `json:"selfDestructKillsMostInGame"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost     int    `json:"gamesLost"`
					GamesPlayed   int    `json:"gamesPlayed"`
					GamesWon      int    `json:"gamesWon"`
					TimePlayed    string `json:"timePlayed"`
					WinPercentage string `json:"winPercentage"`
				} `json:"game"`
				MatchAwards struct {
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous struct {
					TurretsDestroyed int `json:"turretsDestroyed"`
				} `json:"miscellaneous"`
			} `json:"dVa"`
			Doomfist struct {
				Assists interface{} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min  float64 `json:"allDamageDoneAvgPer10Min"`
					DeathsAvgPer10Min         float64 `json:"deathsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min float64 `json:"heroDamageDoneAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame  int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife  int    `json:"allDamageDoneMostInLife"`
					HeroDamageDoneMostInGame int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife int    `json:"heroDamageDoneMostInLife"`
					WeaponAccuracyBestInGame string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					DamageDone     int    `json:"damageDone"`
					Deaths         int    `json:"deaths"`
					HeroDamageDone int    `json:"heroDamageDone"`
					WeaponAccuracy string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					AbilityDamageDone            int     `json:"abilityDamageDone"`
					AbilityDamageDoneAvgPer10Min float64 `json:"abilityDamageDoneAvgPer10Min"`
					AbilityDamageDoneMostInGame  int     `json:"abilityDamageDoneMostInGame"`
					SecondaryFireAccuracy        string  `json:"secondaryFireAccuracy"`
					ShieldsCreated               int     `json:"shieldsCreated"`
					ShieldsCreatedAvgPer10Min    float64 `json:"shieldsCreatedAvgPer10Min"`
					ShieldsCreatedMostInGame     int     `json:"shieldsCreatedMostInGame"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost   int    `json:"gamesLost"`
					GamesPlayed int    `json:"gamesPlayed"`
					TimePlayed  string `json:"timePlayed"`
				} `json:"game"`
				MatchAwards struct {
					Medals       int `json:"medals"`
					MedalsGold   int `json:"medalsGold"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous interface{} `json:"miscellaneous"`
			} `json:"doomfist"`
			Hanzo struct {
				Assists interface{} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
					BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
					CriticalHitsAvgPer10Min      float64 `json:"criticalHitsAvgPer10Min"`
					DeathsAvgPer10Min            float64 `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min      float64 `json:"eliminationsAvgPer10Min"`
					EliminationsPerLife          float64 `json:"eliminationsPerLife"`
					FinalBlowsAvgPer10Min        float64 `json:"finalBlowsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
					ObjectiveKillsAvgPer10Min    float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min     string  `json:"objectiveTimeAvgPer10Min"`
					SoloKillsAvgPer10Min         float64 `json:"soloKillsAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife     int    `json:"allDamageDoneMostInLife"`
					BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
					CriticalHitsMostInGame      int    `json:"criticalHitsMostInGame"`
					CriticalHitsMostInLife      int    `json:"criticalHitsMostInLife"`
					EliminationsMostInGame      int    `json:"eliminationsMostInGame"`
					EliminationsMostInLife      int    `json:"eliminationsMostInLife"`
					FinalBlowsMostInGame        int    `json:"finalBlowsMostInGame"`
					HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife    int    `json:"heroDamageDoneMostInLife"`
					KillsStreakBest             int    `json:"killsStreakBest"`
					ObjectiveKillsMostInGame    int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame     string `json:"objectiveTimeMostInGame"`
					SoloKillsMostInGame         int    `json:"soloKillsMostInGame"`
					WeaponAccuracyBestInGame    string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					BarrierDamageDone    int    `json:"barrierDamageDone"`
					CriticalHits         int    `json:"criticalHits"`
					CriticalHitsAccuracy string `json:"criticalHitsAccuracy"`
					DamageDone           int    `json:"damageDone"`
					Deaths               int    `json:"deaths"`
					Eliminations         int    `json:"eliminations"`
					FinalBlows           int    `json:"finalBlows"`
					HeroDamageDone       int    `json:"heroDamageDone"`
					ObjectiveKills       int    `json:"objectiveKills"`
					ObjectiveTime        string `json:"objectiveTime"`
					QuickMeleeAccuracy   string `json:"quickMeleeAccuracy"`
					SoloKills            int    `json:"soloKills"`
					WeaponAccuracy       string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					DragonstrikeKills            int    `json:"dragonstrikeKills"`
					DragonstrikeKillsAvgPer10Min int    `json:"dragonstrikeKillsAvgPer10Min"`
					DragonstrikeKillsMostInGame  int    `json:"dragonstrikeKillsMostInGame"`
					SecondaryFireAccuracy        string `json:"secondaryFireAccuracy"`
					StormArrowKills              int    `json:"stormArrowKills"`
					StormArrowKillsAvgPer10Min   int    `json:"stormArrowKillsAvgPer10Min"`
					StormArrowKillsMostInGame    int    `json:"stormArrowKillsMostInGame"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost   int    `json:"gamesLost"`
					GamesPlayed int    `json:"gamesPlayed"`
					TimePlayed  string `json:"timePlayed"`
				} `json:"game"`
				MatchAwards struct {
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous struct {
					TurretsDestroyed int `json:"turretsDestroyed"`
				} `json:"miscellaneous"`
			} `json:"hanzo"`
			Roadhog struct {
				Assists struct {
					OffensiveAssists            int     `json:"offensiveAssists"`
					OffensiveAssistsAvgPer10Min float64 `json:"offensiveAssistsAvgPer10Min"`
					OffensiveAssistsMostInGame  int     `json:"offensiveAssistsMostInGame"`
				} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min  float64 `json:"allDamageDoneAvgPer10Min"`
					CriticalHitsAvgPer10Min   float64 `json:"criticalHitsAvgPer10Min"`
					DeathsAvgPer10Min         int     `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min   float64 `json:"eliminationsAvgPer10Min"`
					EliminationsPerLife       int     `json:"eliminationsPerLife"`
					FinalBlowsAvgPer10Min     float64 `json:"finalBlowsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min float64 `json:"heroDamageDoneAvgPer10Min"`
					ObjectiveKillsAvgPer10Min float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min  string  `json:"objectiveTimeAvgPer10Min"`
					SoloKillsAvgPer10Min      int     `json:"soloKillsAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame  int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife  int    `json:"allDamageDoneMostInLife"`
					CriticalHitsMostInGame   int    `json:"criticalHitsMostInGame"`
					CriticalHitsMostInLife   int    `json:"criticalHitsMostInLife"`
					EliminationsMostInGame   int    `json:"eliminationsMostInGame"`
					EliminationsMostInLife   int    `json:"eliminationsMostInLife"`
					FinalBlowsMostInGame     int    `json:"finalBlowsMostInGame"`
					HeroDamageDoneMostInGame int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife int    `json:"heroDamageDoneMostInLife"`
					KillsStreakBest          int    `json:"killsStreakBest"`
					ObjectiveKillsMostInGame int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame  string `json:"objectiveTimeMostInGame"`
					SoloKillsMostInGame      int    `json:"soloKillsMostInGame"`
					WeaponAccuracyBestInGame string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					CriticalHits         int    `json:"criticalHits"`
					CriticalHitsAccuracy string `json:"criticalHitsAccuracy"`
					DamageDone           int    `json:"damageDone"`
					Deaths               int    `json:"deaths"`
					Eliminations         int    `json:"eliminations"`
					FinalBlows           int    `json:"finalBlows"`
					HeroDamageDone       int    `json:"heroDamageDone"`
					ObjectiveKills       int    `json:"objectiveKills"`
					ObjectiveTime        string `json:"objectiveTime"`
					QuickMeleeAccuracy   string `json:"quickMeleeAccuracy"`
					SoloKills            int    `json:"soloKills"`
					WeaponAccuracy       string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					EnemiesHooked            int     `json:"enemiesHooked"`
					EnemiesHookedAvgPer10Min float64 `json:"enemiesHookedAvgPer10Min"`
					EnemiesHookedMostInGame  int     `json:"enemiesHookedMostInGame"`
					HookAccuracy             string  `json:"hookAccuracy"`
					HookAccuracyBestInGame   string  `json:"hookAccuracyBestInGame"`
					HooksAttempted           int     `json:"hooksAttempted"`
					SecondaryFireAccuracy    string  `json:"secondaryFireAccuracy"`
					SelfHealing              int     `json:"selfHealing"`
					SelfHealingAvgPer10Min   float64 `json:"selfHealingAvgPer10Min"`
					SelfHealingMostInGame    int     `json:"selfHealingMostInGame"`
					WholeHogKills            int     `json:"wholeHogKills"`
					WholeHogKillsAvgPer10Min float64 `json:"wholeHogKillsAvgPer10Min"`
					WholeHogKillsMostInGame  int     `json:"wholeHogKillsMostInGame"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost     int    `json:"gamesLost"`
					GamesPlayed   int    `json:"gamesPlayed"`
					GamesWon      int    `json:"gamesWon"`
					TimePlayed    string `json:"timePlayed"`
					WinPercentage string `json:"winPercentage"`
				} `json:"game"`
				MatchAwards struct {
					Medals       int `json:"medals"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous interface{} `json:"miscellaneous"`
			} `json:"roadhog"`
			Tracer struct {
				Assists interface{} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min   float64 `json:"allDamageDoneAvgPer10Min"`
					CriticalHitsAvgPer10Min    float64 `json:"criticalHitsAvgPer10Min"`
					DeathsAvgPer10Min          float64 `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min    float64 `json:"eliminationsAvgPer10Min"`
					EliminationsPerLife        int     `json:"eliminationsPerLife"`
					FinalBlowsAvgPer10Min      float64 `json:"finalBlowsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min  float64 `json:"heroDamageDoneAvgPer10Min"`
					ObjectiveKillsAvgPer10Min  float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min   string  `json:"objectiveTimeAvgPer10Min"`
					TimeSpentOnFireAvgPer10Min string  `json:"timeSpentOnFireAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame   int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife   int    `json:"allDamageDoneMostInLife"`
					CriticalHitsMostInGame    int    `json:"criticalHitsMostInGame"`
					CriticalHitsMostInLife    int    `json:"criticalHitsMostInLife"`
					EliminationsMostInGame    int    `json:"eliminationsMostInGame"`
					EliminationsMostInLife    int    `json:"eliminationsMostInLife"`
					FinalBlowsMostInGame      int    `json:"finalBlowsMostInGame"`
					HeroDamageDoneMostInGame  int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife  int    `json:"heroDamageDoneMostInLife"`
					KillsStreakBest           int    `json:"killsStreakBest"`
					ObjectiveKillsMostInGame  int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame   string `json:"objectiveTimeMostInGame"`
					TimeSpentOnFireMostInGame string `json:"timeSpentOnFireMostInGame"`
					WeaponAccuracyBestInGame  string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					CriticalHits         int    `json:"criticalHits"`
					CriticalHitsAccuracy string `json:"criticalHitsAccuracy"`
					DamageDone           int    `json:"damageDone"`
					Deaths               int    `json:"deaths"`
					Eliminations         int    `json:"eliminations"`
					FinalBlows           int    `json:"finalBlows"`
					HeroDamageDone       int    `json:"heroDamageDone"`
					ObjectiveKills       int    `json:"objectiveKills"`
					ObjectiveTime        string `json:"objectiveTime"`
					QuickMeleeAccuracy   string `json:"quickMeleeAccuracy"`
					TimeSpentOnFire      string `json:"timeSpentOnFire"`
					WeaponAccuracy       string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					SecondaryFireAccuracy string `json:"secondaryFireAccuracy"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost     int    `json:"gamesLost"`
					GamesPlayed   int    `json:"gamesPlayed"`
					GamesWon      int    `json:"gamesWon"`
					TimePlayed    string `json:"timePlayed"`
					WinPercentage string `json:"winPercentage"`
				} `json:"game"`
				MatchAwards struct {
					Medals       int `json:"medals"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous interface{} `json:"miscellaneous"`
			} `json:"tracer"`
			Widowmaker struct {
				Assists interface{} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
					BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
					DeathsAvgPer10Min            float64 `json:"deathsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife     int    `json:"allDamageDoneMostInLife"`
					BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
					HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife    int    `json:"heroDamageDoneMostInLife"`
					WeaponAccuracyBestInGame    string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					BarrierDamageDone int    `json:"barrierDamageDone"`
					DamageDone        int    `json:"damageDone"`
					Deaths            int    `json:"deaths"`
					HeroDamageDone    int    `json:"heroDamageDone"`
					WeaponAccuracy    string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					ScopedAccuracy           string `json:"scopedAccuracy"`
					ScopedAccuracyBestInGame string `json:"scopedAccuracyBestInGame"`
					SecondaryFireAccuracy    string `json:"secondaryFireAccuracy"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost   int    `json:"gamesLost"`
					GamesPlayed int    `json:"gamesPlayed"`
					TimePlayed  string `json:"timePlayed"`
				} `json:"game"`
				MatchAwards struct {
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous interface{} `json:"miscellaneous"`
			} `json:"widowmaker"`
			Winston struct {
				Assists interface{} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
					BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
					DeathsAvgPer10Min            float64 `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min      float64 `json:"eliminationsAvgPer10Min"`
					EliminationsPerLife          float64 `json:"eliminationsPerLife"`
					FinalBlowsAvgPer10Min        float64 `json:"finalBlowsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
					ObjectiveKillsAvgPer10Min    float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min     string  `json:"objectiveTimeAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife     int    `json:"allDamageDoneMostInLife"`
					BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
					EliminationsMostInGame      int    `json:"eliminationsMostInGame"`
					EliminationsMostInLife      int    `json:"eliminationsMostInLife"`
					FinalBlowsMostInGame        int    `json:"finalBlowsMostInGame"`
					HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife    int    `json:"heroDamageDoneMostInLife"`
					KillsStreakBest             int    `json:"killsStreakBest"`
					ObjectiveKillsMostInGame    int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame     string `json:"objectiveTimeMostInGame"`
				} `json:"best"`
				Combat struct {
					BarrierDamageDone  int    `json:"barrierDamageDone"`
					DamageDone         int    `json:"damageDone"`
					Deaths             int    `json:"deaths"`
					Eliminations       int    `json:"eliminations"`
					FinalBlows         int    `json:"finalBlows"`
					HeroDamageDone     int    `json:"heroDamageDone"`
					ObjectiveKills     int    `json:"objectiveKills"`
					ObjectiveTime      string `json:"objectiveTime"`
					QuickMeleeAccuracy string `json:"quickMeleeAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					DamageBlocked                 int     `json:"damageBlocked"`
					DamageBlockedAvgPer10Min      float64 `json:"damageBlockedAvgPer10Min"`
					DamageBlockedMostInGame       int     `json:"damageBlockedMostInGame"`
					JumpKills                     int     `json:"jumpKills"`
					JumpPackKills                 int     `json:"jumpPackKills"`
					JumpPackKillsAvgPer10Min      float64 `json:"jumpPackKillsAvgPer10Min"`
					JumpPackKillsMostInGame       int     `json:"jumpPackKillsMostInGame"`
					MeleeKills                    int     `json:"meleeKills"`
					MeleeKillsAvgPer10Min         float64 `json:"meleeKillsAvgPer10Min"`
					MeleeKillsMostInGame          int     `json:"meleeKillsMostInGame"`
					PlayersKnockedBack            int     `json:"playersKnockedBack"`
					PlayersKnockedBackAvgPer10Min float64 `json:"playersKnockedBackAvgPer10Min"`
					PlayersKnockedBackMostInGame  int     `json:"playersKnockedBackMostInGame"`
					PrimalRageKills               int     `json:"primalRageKills"`
					PrimalRageKillsAvgPer10Min    int     `json:"primalRageKillsAvgPer10Min"`
					PrimalRageKillsMostInGame     int     `json:"primalRageKillsMostInGame"`
					PrimalRageMeleeAccuracy       string  `json:"primalRageMeleeAccuracy"`
					TeslaCannonAccuracy           string  `json:"teslaCannonAccuracy"`
					WeaponKills                   int     `json:"weaponKills"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost   int    `json:"gamesLost"`
					GamesPlayed int    `json:"gamesPlayed"`
					TimePlayed  string `json:"timePlayed"`
				} `json:"game"`
				MatchAwards struct {
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsGold   int `json:"medalsGold"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous interface{} `json:"miscellaneous"`
			} `json:"winston"`
			WreckingBall struct {
				Assists interface{} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min  float64 `json:"allDamageDoneAvgPer10Min"`
					HeroDamageDoneAvgPer10Min float64 `json:"heroDamageDoneAvgPer10Min"`
					ObjectiveTimeAvgPer10Min  string  `json:"objectiveTimeAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame  int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife  int    `json:"allDamageDoneMostInLife"`
					HeroDamageDoneMostInGame int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife int    `json:"heroDamageDoneMostInLife"`
					ObjectiveTimeMostInGame  string `json:"objectiveTimeMostInGame"`
					WeaponAccuracyBestInGame string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					DamageDone     int    `json:"damageDone"`
					HeroDamageDone int    `json:"heroDamageDone"`
					ObjectiveTime  string `json:"objectiveTime"`
					WeaponAccuracy string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					PlayersKnockedBack            int     `json:"playersKnockedBack"`
					PlayersKnockedBackAvgPer10Min float64 `json:"playersKnockedBackAvgPer10Min"`
					PlayersKnockedBackMostInGame  int     `json:"playersKnockedBackMostInGame"`
					SecondaryFireAccuracy         string  `json:"secondaryFireAccuracy"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost   int    `json:"gamesLost"`
					GamesPlayed int    `json:"gamesPlayed"`
					TimePlayed  string `json:"timePlayed"`
				} `json:"game"`
				MatchAwards struct {
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsGold   int `json:"medalsGold"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous interface{} `json:"miscellaneous"`
			} `json:"wreckingBall"`
			Zenyatta struct {
				Assists struct {
					DefensiveAssists            int     `json:"defensiveAssists"`
					DefensiveAssistsAvgPer10Min float64 `json:"defensiveAssistsAvgPer10Min"`
					DefensiveAssistsMostInGame  int     `json:"defensiveAssistsMostInGame"`
					HealingDone                 int     `json:"healingDone"`
					HealingDoneAvgPer10Min      float64 `json:"healingDoneAvgPer10Min"`
					HealingDoneMostInGame       int     `json:"healingDoneMostInGame"`
					OffensiveAssists            int     `json:"offensiveAssists"`
					OffensiveAssistsAvgPer10Min float64 `json:"offensiveAssistsAvgPer10Min"`
					OffensiveAssistsMostInGame  int     `json:"offensiveAssistsMostInGame"`
				} `json:"assists"`
				Average struct {
					AllDamageDoneAvgPer10Min     float64 `json:"allDamageDoneAvgPer10Min"`
					BarrierDamageDoneAvgPer10Min float64 `json:"barrierDamageDoneAvgPer10Min"`
					CriticalHitsAvgPer10Min      float64 `json:"criticalHitsAvgPer10Min"`
					DeathsAvgPer10Min            float64 `json:"deathsAvgPer10Min"`
					EliminationsAvgPer10Min      float64 `json:"eliminationsAvgPer10Min"`
					EliminationsPerLife          float64 `json:"eliminationsPerLife"`
					FinalBlowsAvgPer10Min        float64 `json:"finalBlowsAvgPer10Min"`
					HeroDamageDoneAvgPer10Min    float64 `json:"heroDamageDoneAvgPer10Min"`
					ObjectiveKillsAvgPer10Min    float64 `json:"objectiveKillsAvgPer10Min"`
					ObjectiveTimeAvgPer10Min     string  `json:"objectiveTimeAvgPer10Min"`
					SoloKillsAvgPer10Min         int     `json:"soloKillsAvgPer10Min"`
					TimeSpentOnFireAvgPer10Min   string  `json:"timeSpentOnFireAvgPer10Min"`
				} `json:"average"`
				Best struct {
					AllDamageDoneMostInGame     int    `json:"allDamageDoneMostInGame"`
					AllDamageDoneMostInLife     int    `json:"allDamageDoneMostInLife"`
					BarrierDamageDoneMostInGame int    `json:"barrierDamageDoneMostInGame"`
					CriticalHitsMostInGame      int    `json:"criticalHitsMostInGame"`
					CriticalHitsMostInLife      int    `json:"criticalHitsMostInLife"`
					EliminationsMostInGame      int    `json:"eliminationsMostInGame"`
					EliminationsMostInLife      int    `json:"eliminationsMostInLife"`
					FinalBlowsMostInGame        int    `json:"finalBlowsMostInGame"`
					HeroDamageDoneMostInGame    int    `json:"heroDamageDoneMostInGame"`
					HeroDamageDoneMostInLife    int    `json:"heroDamageDoneMostInLife"`
					KillsStreakBest             int    `json:"killsStreakBest"`
					ObjectiveKillsMostInGame    int    `json:"objectiveKillsMostInGame"`
					ObjectiveTimeMostInGame     string `json:"objectiveTimeMostInGame"`
					SoloKillsMostInGame         int    `json:"soloKillsMostInGame"`
					TimeSpentOnFireMostInGame   string `json:"timeSpentOnFireMostInGame"`
					WeaponAccuracyBestInGame    string `json:"weaponAccuracyBestInGame"`
				} `json:"best"`
				Combat struct {
					BarrierDamageDone    int    `json:"barrierDamageDone"`
					CriticalHits         int    `json:"criticalHits"`
					CriticalHitsAccuracy string `json:"criticalHitsAccuracy"`
					DamageDone           int    `json:"damageDone"`
					Deaths               int    `json:"deaths"`
					Eliminations         int    `json:"eliminations"`
					FinalBlows           int    `json:"finalBlows"`
					HeroDamageDone       int    `json:"heroDamageDone"`
					ObjectiveKills       int    `json:"objectiveKills"`
					ObjectiveTime        string `json:"objectiveTime"`
					QuickMeleeAccuracy   string `json:"quickMeleeAccuracy"`
					SoloKills            int    `json:"soloKills"`
					TimeSpentOnFire      string `json:"timeSpentOnFire"`
					WeaponAccuracy       string `json:"weaponAccuracy"`
				} `json:"combat"`
				Deaths       interface{} `json:"deaths"`
				HeroSpecific struct {
					SecondaryFireAccuracy    string  `json:"secondaryFireAccuracy"`
					SelfHealing              int     `json:"selfHealing"`
					SelfHealingAvgPer10Min   float64 `json:"selfHealingAvgPer10Min"`
					SelfHealingMostInGame    int     `json:"selfHealingMostInGame"`
					TranscendenceHealing     int     `json:"transcendenceHealing"`
					TranscendenceHealingBest int     `json:"transcendenceHealingBest"`
				} `json:"heroSpecific"`
				Game struct {
					GamesLost     int    `json:"gamesLost"`
					GamesPlayed   int    `json:"gamesPlayed"`
					GamesWon      int    `json:"gamesWon"`
					TimePlayed    string `json:"timePlayed"`
					WinPercentage string `json:"winPercentage"`
				} `json:"game"`
				MatchAwards struct {
					Cards        int `json:"cards"`
					Medals       int `json:"medals"`
					MedalsBronze int `json:"medalsBronze"`
					MedalsGold   int `json:"medalsGold"`
					MedalsSilver int `json:"medalsSilver"`
				} `json:"matchAwards"`
				Miscellaneous interface{} `json:"miscellaneous"`
			} `json:"zenyatta"`
		} `json:"careerStats"`
		Games struct {
			Played int `json:"played"`
			Won    int `json:"won"`
		} `json:"games"`
		TopHeroes struct {
			Ana struct {
				TimePlayed          string  `json:"timePlayed"`
				TimePlayedInSeconds int     `json:"timePlayedInSeconds"`
				GamesWon            int     `json:"gamesWon"`
				WinPercentage       int     `json:"winPercentage"`
				WeaponAccuracy      int     `json:"weaponAccuracy"`
				EliminationsPerLife float64 `json:"eliminationsPerLife"`
				MultiKillBest       int     `json:"multiKillBest"`
				ObjectiveKills      int     `json:"objectiveKills"`
			} `json:"ana"`
			Baptiste struct {
				TimePlayed          string `json:"timePlayed"`
				TimePlayedInSeconds int    `json:"timePlayedInSeconds"`
				GamesWon            int    `json:"gamesWon"`
				WinPercentage       int    `json:"winPercentage"`
				WeaponAccuracy      int    `json:"weaponAccuracy"`
				EliminationsPerLife int    `json:"eliminationsPerLife"`
				MultiKillBest       int    `json:"multiKillBest"`
				ObjectiveKills      int    `json:"objectiveKills"`
			} `json:"baptiste"`
			DVa struct {
				TimePlayed          string  `json:"timePlayed"`
				TimePlayedInSeconds int     `json:"timePlayedInSeconds"`
				GamesWon            int     `json:"gamesWon"`
				WinPercentage       int     `json:"winPercentage"`
				WeaponAccuracy      int     `json:"weaponAccuracy"`
				EliminationsPerLife float64 `json:"eliminationsPerLife"`
				MultiKillBest       int     `json:"multiKillBest"`
				ObjectiveKills      int     `json:"objectiveKills"`
			} `json:"dVa"`
			Doomfist struct {
				TimePlayed          string `json:"timePlayed"`
				TimePlayedInSeconds int    `json:"timePlayedInSeconds"`
				GamesWon            int    `json:"gamesWon"`
				WinPercentage       int    `json:"winPercentage"`
				WeaponAccuracy      int    `json:"weaponAccuracy"`
				EliminationsPerLife int    `json:"eliminationsPerLife"`
				MultiKillBest       int    `json:"multiKillBest"`
				ObjectiveKills      int    `json:"objectiveKills"`
			} `json:"doomfist"`
			Hanzo struct {
				TimePlayed          string  `json:"timePlayed"`
				TimePlayedInSeconds int     `json:"timePlayedInSeconds"`
				GamesWon            int     `json:"gamesWon"`
				WinPercentage       int     `json:"winPercentage"`
				WeaponAccuracy      int     `json:"weaponAccuracy"`
				EliminationsPerLife float64 `json:"eliminationsPerLife"`
				MultiKillBest       int     `json:"multiKillBest"`
				ObjectiveKills      int     `json:"objectiveKills"`
			} `json:"hanzo"`
			Roadhog struct {
				TimePlayed          string `json:"timePlayed"`
				TimePlayedInSeconds int    `json:"timePlayedInSeconds"`
				GamesWon            int    `json:"gamesWon"`
				WinPercentage       int    `json:"winPercentage"`
				WeaponAccuracy      int    `json:"weaponAccuracy"`
				EliminationsPerLife int    `json:"eliminationsPerLife"`
				MultiKillBest       int    `json:"multiKillBest"`
				ObjectiveKills      int    `json:"objectiveKills"`
			} `json:"roadhog"`
			Tracer struct {
				TimePlayed          string `json:"timePlayed"`
				TimePlayedInSeconds int    `json:"timePlayedInSeconds"`
				GamesWon            int    `json:"gamesWon"`
				WinPercentage       int    `json:"winPercentage"`
				WeaponAccuracy      int    `json:"weaponAccuracy"`
				EliminationsPerLife int    `json:"eliminationsPerLife"`
				MultiKillBest       int    `json:"multiKillBest"`
				ObjectiveKills      int    `json:"objectiveKills"`
			} `json:"tracer"`
			Widowmaker struct {
				TimePlayed          string `json:"timePlayed"`
				TimePlayedInSeconds int    `json:"timePlayedInSeconds"`
				GamesWon            int    `json:"gamesWon"`
				WinPercentage       int    `json:"winPercentage"`
				WeaponAccuracy      int    `json:"weaponAccuracy"`
				EliminationsPerLife int    `json:"eliminationsPerLife"`
				MultiKillBest       int    `json:"multiKillBest"`
				ObjectiveKills      int    `json:"objectiveKills"`
			} `json:"widowmaker"`
			Winston struct {
				TimePlayed          string  `json:"timePlayed"`
				TimePlayedInSeconds int     `json:"timePlayedInSeconds"`
				GamesWon            int     `json:"gamesWon"`
				WinPercentage       int     `json:"winPercentage"`
				WeaponAccuracy      int     `json:"weaponAccuracy"`
				EliminationsPerLife float64 `json:"eliminationsPerLife"`
				MultiKillBest       int     `json:"multiKillBest"`
				ObjectiveKills      int     `json:"objectiveKills"`
			} `json:"winston"`
			WreckingBall struct {
				TimePlayed          string `json:"timePlayed"`
				TimePlayedInSeconds int    `json:"timePlayedInSeconds"`
				GamesWon            int    `json:"gamesWon"`
				WinPercentage       int    `json:"winPercentage"`
				WeaponAccuracy      int    `json:"weaponAccuracy"`
				EliminationsPerLife int    `json:"eliminationsPerLife"`
				MultiKillBest       int    `json:"multiKillBest"`
				ObjectiveKills      int    `json:"objectiveKills"`
			} `json:"wreckingBall"`
			Zenyatta struct {
				TimePlayed          string  `json:"timePlayed"`
				TimePlayedInSeconds int     `json:"timePlayedInSeconds"`
				GamesWon            int     `json:"gamesWon"`
				WinPercentage       int     `json:"winPercentage"`
				WeaponAccuracy      int     `json:"weaponAccuracy"`
				EliminationsPerLife float64 `json:"eliminationsPerLife"`
				MultiKillBest       int     `json:"multiKillBest"`
				ObjectiveKills      int     `json:"objectiveKills"`
			} `json:"zenyatta"`
		} `json:"topHeroes"`
	} `json:"competitiveStats"`
	Endorsement     int    `json:"endorsement"`
	EndorsementIcon string `json:"endorsementIcon"`
	GamesWon        int    `json:"gamesWon"`
	Icon            string `json:"icon"`
	Level           int    `json:"level"`
	LevelIcon       string `json:"levelIcon"`
	Name            string `json:"name"`
	Prestige        int    `json:"prestige"`
	PrestigeIcon    string `json:"prestigeIcon"`
	Private         bool   `json:"private"`
	Rating          int    `json:"rating"`
	RatingIcon      string `json:"ratingIcon"`
}

func getPlayerStats(battletag string, platform string, region string) (*OWPlayer, error) {

	//Prepare endpoint url part
	if region != "" {
		region = "/" + region
	}
	url := fmt.Sprintf("https://ow-api.com/v1/stats/%v%v/%v/profile", platform, region, battletag)
	requ, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.New("An error while retrieving data from the Overwatch stats api occured.\n" + err.Error())
	}

	resp, err := Client.Do(requ)
	if err != nil {
		return nil, errors.New("An error while retrieving data from the Overwatch stats api occured.\n" + err.Error())
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("An error while reading the response from the Overwatch API, player request.\n" + err.Error())
	}

	var owPlayerStats OWPlayer
	err = json.Unmarshal(bytes, &owPlayerStats)
	if err != nil {
		return nil, err
	}

	return &owPlayerStats, nil
}
