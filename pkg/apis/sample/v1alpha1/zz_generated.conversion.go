// +build !ignore_autogenerated

/*

Copyright 2018 This Project Authors.

Author:  seanchann <seanchann@foxmail.com>

See docs/ for more information about the  project.

*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	sample "github.com/seanchann/sample-apiserver/pkg/apis/sample"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Test_To_sample_Test,
		Convert_sample_Test_To_v1alpha1_Test,
		Convert_v1alpha1_TestList_To_sample_TestList,
		Convert_sample_TestList_To_v1alpha1_TestList,
		Convert_v1alpha1_TestSpec_To_sample_TestSpec,
		Convert_sample_TestSpec_To_v1alpha1_TestSpec,
		Convert_v1alpha1_User_To_sample_User,
		Convert_sample_User_To_v1alpha1_User,
		Convert_v1alpha1_UserList_To_sample_UserList,
		Convert_sample_UserList_To_v1alpha1_UserList,
		Convert_v1alpha1_UserSpec_To_sample_UserSpec,
		Convert_sample_UserSpec_To_v1alpha1_UserSpec,
	)
}

func autoConvert_v1alpha1_Test_To_sample_Test(in *Test, out *sample.Test, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_TestSpec_To_sample_TestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Test_To_sample_Test is an autogenerated conversion function.
func Convert_v1alpha1_Test_To_sample_Test(in *Test, out *sample.Test, s conversion.Scope) error {
	return autoConvert_v1alpha1_Test_To_sample_Test(in, out, s)
}

func autoConvert_sample_Test_To_v1alpha1_Test(in *sample.Test, out *Test, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_sample_TestSpec_To_v1alpha1_TestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_sample_Test_To_v1alpha1_Test is an autogenerated conversion function.
func Convert_sample_Test_To_v1alpha1_Test(in *sample.Test, out *Test, s conversion.Scope) error {
	return autoConvert_sample_Test_To_v1alpha1_Test(in, out, s)
}

func autoConvert_v1alpha1_TestList_To_sample_TestList(in *TestList, out *sample.TestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]sample.Test)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_TestList_To_sample_TestList is an autogenerated conversion function.
func Convert_v1alpha1_TestList_To_sample_TestList(in *TestList, out *sample.TestList, s conversion.Scope) error {
	return autoConvert_v1alpha1_TestList_To_sample_TestList(in, out, s)
}

func autoConvert_sample_TestList_To_v1alpha1_TestList(in *sample.TestList, out *TestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Test)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_sample_TestList_To_v1alpha1_TestList is an autogenerated conversion function.
func Convert_sample_TestList_To_v1alpha1_TestList(in *sample.TestList, out *TestList, s conversion.Scope) error {
	return autoConvert_sample_TestList_To_v1alpha1_TestList(in, out, s)
}

func autoConvert_v1alpha1_TestSpec_To_sample_TestSpec(in *TestSpec, out *sample.TestSpec, s conversion.Scope) error {
	out.Family = in.Family
	return nil
}

// Convert_v1alpha1_TestSpec_To_sample_TestSpec is an autogenerated conversion function.
func Convert_v1alpha1_TestSpec_To_sample_TestSpec(in *TestSpec, out *sample.TestSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_TestSpec_To_sample_TestSpec(in, out, s)
}

func autoConvert_sample_TestSpec_To_v1alpha1_TestSpec(in *sample.TestSpec, out *TestSpec, s conversion.Scope) error {
	out.Family = in.Family
	return nil
}

// Convert_sample_TestSpec_To_v1alpha1_TestSpec is an autogenerated conversion function.
func Convert_sample_TestSpec_To_v1alpha1_TestSpec(in *sample.TestSpec, out *TestSpec, s conversion.Scope) error {
	return autoConvert_sample_TestSpec_To_v1alpha1_TestSpec(in, out, s)
}

func autoConvert_v1alpha1_User_To_sample_User(in *User, out *sample.User, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_UserSpec_To_sample_UserSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_User_To_sample_User is an autogenerated conversion function.
func Convert_v1alpha1_User_To_sample_User(in *User, out *sample.User, s conversion.Scope) error {
	return autoConvert_v1alpha1_User_To_sample_User(in, out, s)
}

func autoConvert_sample_User_To_v1alpha1_User(in *sample.User, out *User, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_sample_UserSpec_To_v1alpha1_UserSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_sample_User_To_v1alpha1_User is an autogenerated conversion function.
func Convert_sample_User_To_v1alpha1_User(in *sample.User, out *User, s conversion.Scope) error {
	return autoConvert_sample_User_To_v1alpha1_User(in, out, s)
}

func autoConvert_v1alpha1_UserList_To_sample_UserList(in *UserList, out *sample.UserList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]sample.User)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_UserList_To_sample_UserList is an autogenerated conversion function.
func Convert_v1alpha1_UserList_To_sample_UserList(in *UserList, out *sample.UserList, s conversion.Scope) error {
	return autoConvert_v1alpha1_UserList_To_sample_UserList(in, out, s)
}

func autoConvert_sample_UserList_To_v1alpha1_UserList(in *sample.UserList, out *UserList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]User)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_sample_UserList_To_v1alpha1_UserList is an autogenerated conversion function.
func Convert_sample_UserList_To_v1alpha1_UserList(in *sample.UserList, out *UserList, s conversion.Scope) error {
	return autoConvert_sample_UserList_To_v1alpha1_UserList(in, out, s)
}

func autoConvert_v1alpha1_UserSpec_To_sample_UserSpec(in *UserSpec, out *sample.UserSpec, s conversion.Scope) error {
	out.ID = in.ID
	out.Passwd = in.Passwd
	out.Email = in.Email
	out.Name = in.Name
	out.EmailVerify = in.EmailVerify
	out.Status = in.Status
	out.RawObj = *(*[]byte)(unsafe.Pointer(&in.RawObj))
	return nil
}

// Convert_v1alpha1_UserSpec_To_sample_UserSpec is an autogenerated conversion function.
func Convert_v1alpha1_UserSpec_To_sample_UserSpec(in *UserSpec, out *sample.UserSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_UserSpec_To_sample_UserSpec(in, out, s)
}

func autoConvert_sample_UserSpec_To_v1alpha1_UserSpec(in *sample.UserSpec, out *UserSpec, s conversion.Scope) error {
	out.ID = in.ID
	out.Passwd = in.Passwd
	out.Email = in.Email
	out.Name = in.Name
	out.EmailVerify = in.EmailVerify
	out.Status = in.Status
	out.RawObj = *(*[]byte)(unsafe.Pointer(&in.RawObj))
	return nil
}

// Convert_sample_UserSpec_To_v1alpha1_UserSpec is an autogenerated conversion function.
func Convert_sample_UserSpec_To_v1alpha1_UserSpec(in *sample.UserSpec, out *UserSpec, s conversion.Scope) error {
	return autoConvert_sample_UserSpec_To_v1alpha1_UserSpec(in, out, s)
}
