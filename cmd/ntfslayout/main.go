package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gentlemanautomaton/ntfs"
	"github.com/gentlemanautomaton/ntfs/mspart"
	"github.com/rekby/gpt"
)

const sectorSize = 512 // TODO: Try to detect this automatically

func main() {
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		fmt.Printf("usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Printf("Opening \"%s\"\n", path)

	// Open the raw file
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Unable to open \"%s\": %s\n", path, err)
		os.Exit(1)
	}
	defer f.Close()
	defer fmt.Printf("Closing \"%s\"\n", path)

	// Seek past the protective MBR to the second sector (assuming sector size for now)
	f.Seek(sectorSize, io.SeekStart)

	// Read the GPT partition table
	table, err := gpt.ReadTable(f, sectorSize)
	if err != nil {
		fmt.Printf("Unable to read GPT partition table: %s\n", err)
		return
	}

	fmt.Printf("--------\nDisk %v\n--------\n", table.Header.DiskGUID)

	for _, part := range table.Partitions {
		if part.IsEmpty() {
			continue
		}

		fmt.Printf("Partition %s: %s Range[%d:%d]\n", part.Name(), part.Type, part.FirstLBA, part.LastLBA)
		if part.Type != mspart.BasicData {
			continue
		}

		start := int64(part.FirstLBA * table.SectorSize)
		end := int64(part.LastLBA*table.SectorSize + table.SectorSize)
		fmt.Printf("  Partition Start:              %d\n", start)
		fmt.Printf("  Partition End:                %d\n", end)
		fmt.Printf("  Partition Size:               %d\n", end-start)

		section := io.NewSectionReader(f, start, end-start)

		r, err := ntfs.NewReader(section)
		if err != nil {
			fmt.Printf("  Unable to read basic data volume NTFS volume boot record: %v\n", err)
			continue
		}

		vbr := r.BootRecord()
		fmt.Printf("  BytesPerSector:               %d\n", vbr.BytesPerSector)
		fmt.Printf("  SectorsPerCluster:            %d\n", vbr.SectorsPerCluster)
		fmt.Printf("  ClusterSize:                  %d\n", vbr.ClusterSize())
		fmt.Printf("  MediaDescriptor:              %d\n", vbr.MediaDescriptor)
		fmt.Printf("  TotalSectors:                 %d\n", vbr.TotalSectors)
		fmt.Printf("  MFT:                          %d\n", vbr.MFT)
		fmt.Printf("  MFTMirror:                    %d\n", vbr.MFTMirror)
		fmt.Printf("  ClustersPerFileRecordSegment: %-6d (%d bytes)\n", vbr.ClustersPerFileRecordSegment, vbr.FileRecordSize())
		fmt.Printf("  ClustersPerIndexBlock:        %-6d (%d bytes)\n", vbr.ClustersPerIndexBlock, vbr.IndexBlockSize())
		fmt.Printf("  VolumeSerialNumber:           %d\n", vbr.VolumeSerialNumber)
		fmt.Printf("  Checksum:                     %d\n", vbr.Checksum)
	}
}
