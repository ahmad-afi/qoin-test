package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	player struct {
		totalPoint int
		totalDice  int
	}
)

func rollTheDice() int {
	rand.Seed(time.Now().UnixNano())
	// karena random dimulai dari 0
	return rand.Intn(6) + 1
}

func GameDadu(numPlayer, numDice int) {

	cpNumPlayer := numPlayer
	mapPlayer := make(map[int]player)

	// Mengeset nilai awal
	for i := 0; i < numPlayer; i++ {
		mapPlayer[i] = player{
			totalPoint: 0,
			totalDice:  numDice,
		}
	}

	// breakLoop := true
	ronde := 1
	// Mulai Permainan
	for {
		tempPoint := make(map[int][]int)

		// Hentikan looping
		if cpNumPlayer == 1 {
			break
		}
		fmt.Println("Kocokan ke-", ronde)

		// Setiap pemain mengocok

		// !!!!!!!!!!!!! error numPlayer kebalik
		for i := 0; i < numPlayer; i++ {
			jumlahKocokan := mapPlayer[i].totalDice

			if jumlahKocokan < 1 {
				continue
			}

			// Looping untuk mengocok
			for j := 0; j < jumlahKocokan; j++ {
				tempPoint[i] = append(tempPoint[i], rollTheDice())
			}
			fmt.Printf("Hasil kocokan permain %d adalah %#v \n", i+1, tempPoint[i])
		}

		// Evaluasi nilai

		// Looping player
		for i := 0; i < numPlayer; i++ {

			// Looping hasil kocokan
			for j := 0; j < len(tempPoint[i]); j++ {
				// Mencopy nilai dari map
				pointerMap := mapPlayer[i]

				pointerValDice := tempPoint[i][j]

				// Mengecek nilai dadu
				if pointerValDice == 6 {
					pointerMap.totalDice -= 1
					pointerMap.totalPoint += 6

					// Cek kalau nilainya sudah 0, kurani total pemain
					if pointerMap.totalDice < 1 {
						cpNumPlayer -= 1
					}

				} else if pointerValDice == 1 {
					// Kalau i lebih besar dari jumlah pemain
					if i+1 >= numPlayer {

						// cek apakah masih bermain, dimulai dari player 0
						idPlayer := 0
						for {
							if mapPlayer[idPlayer].totalDice > 0 {
								// Memberikan nilai ke player selnjutnya
								pointerMapPlayer := mapPlayer[idPlayer]
								pointerMapPlayer.totalPoint += 1

								// Mengeset kembali nilai map
								mapPlayer[idPlayer] = pointerMapPlayer
								break
							}
							idPlayer++
						}

					} else {
						// Memberikan poin ke player selanjutnya
						pointerMapPlayer := mapPlayer[i+1]
						pointerMapPlayer.totalPoint += pointerValDice

						// Mengeset kembali nilai map
						mapPlayer[i+1] = pointerMapPlayer
					}
				} else {
					pointerMap.totalPoint += pointerValDice
				}

				// Mengeset nilai nya ke map
				mapPlayer[i] = pointerMap
			}
		}

		ronde++

		fmt.Println("Hasil sementara")
		for i := 0; i < numPlayer; i++ {
			fmt.Printf("Pemain %d memiliki poin %d dengan total dadu %d \n", i+1, mapPlayer[i].totalPoint, mapPlayer[i].totalDice)
		}
		fmt.Println()
	}

	var pemenang, poin int
	var pemainTerakhir int

	for i := 0; i < numPlayer; i++ {
		if mapPlayer[i].totalDice > 0 {
			pemainTerakhir = i
		}

		if mapPlayer[i].totalPoint > poin {
			poin = mapPlayer[i].totalPoint
			pemenang = i
		}
	}

	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu. \n", pemainTerakhir)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya yaitu %d. \n", pemenang, poin)
}
