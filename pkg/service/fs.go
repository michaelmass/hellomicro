package service

import (
	"context"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	api "github.com/michaelmass/hellomicro/api/proto"
	"github.com/pkg/errors"
	empty "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *Service) ListFiles(ctx context.Context, request *api.ListFilesReq) (*api.ListFilesRes, error) {
	fileInfos := &api.ListFilesRes{
		FileInfos: []*api.FileInfo{},
	}

	files, err := filepath.Glob(request.Path + "/*")

	if err != nil {
		return nil, errors.Wrap(err, "unable to list files")
	}

	for _, file := range files {
		fileInfo, err := os.Stat(file)

		if err != nil {
			return nil, errors.Wrapf(err, "unable get info of file %s", file)
		}

		modTime := timestamppb.New(fileInfo.ModTime())

		fileInfos.FileInfos = append(fileInfos.FileInfos, &api.FileInfo{
			Path:    file,
			Size:    fileInfo.Size(),
			ModTime: modTime,
			Mode:    fileInfo.Mode().String(),
		})
	}

	return fileInfos, nil
}

func (service *Service) WriteFile(ctx context.Context, request *api.WriteFileReq) (*empty.Empty, error) {
	mode, err := strconv.ParseUint(request.Mode, 8, 32)

	if err != nil {
		return nil, errors.Wrapf(err, "unable to convert file mode from string to uint32 %s", request.Mode)
	}

	err = ioutil.WriteFile(request.Path, request.Content, fs.FileMode(mode))

	if err != nil {
		return nil, errors.Wrapf(err, "unable to write file %s", request.Path)
	}

	return e, nil
}

func (service *Service) ReadFile(ctx context.Context, request *api.ReadFileReq) (*api.ReadFileRes, error) {
	return &api.ReadFileRes{}, nil
}

func (service *Service) DeleteFile(ctx context.Context, request *api.DeleteFileReq) (*api.DeleteFileRes, error) {
	return &api.DeleteFileRes{}, nil
}
