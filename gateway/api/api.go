package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func GeneralGetPageLimit(r *http.Request, pageDefault int, limitDefault int) (int, int, error) {
	// 注：本函数仅在对page和limit需要做一些操作时使用，正常情况下采用r.URL.RawQuery直接转发
	// 获取查询参数
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	// 设置默认值
	var pageInt, limitInt int
	var err error
	if page == "" {
		pageInt = pageDefault
	} else {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return 0, 0, err
		}
	}
	if limit == "" {
		limitInt = limitDefault
	} else {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, err
		}
	}
	return pageInt, limitInt, nil
}

func GeneralGetResponse(w http.ResponseWriter, code int, res []byte, e error, errorMessage string) {
	if e != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// 返回响应
	fmt.Fprint(w, string(res))
	return
}

func GeneralPostResponse(w http.ResponseWriter, code int, res []byte, e error, errorMessage string) {
	if e != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	// 将接口的返回原样返回出去
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
	return
}

func GeneralPutResponse(w http.ResponseWriter, code int, res []byte, e error, errorMessage string) {
	if e != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	// 将接口的返回原样返回出去
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
	return
}

func GeneralDeleteResponse(w http.ResponseWriter, code int, res []byte, e error, errorMessage string) {
	if e != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// 返回响应
	fmt.Fprint(w, string(res))
	return
}

func GeneralPatchResponse(w http.ResponseWriter, code int, res []byte, e error, errorMessage string) {
	if e != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	// 将接口的返回原样返回出去
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
	return
}
