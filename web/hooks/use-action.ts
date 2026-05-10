import { ApiError, BaseResponse, PaginatedMeta } from "@/types/base-response";
import { useActionState, useEffect, useRef } from "react";

type ActionOptions<TData> = {
  // Pass both data and meta to onSuccess for convenience
  onSuccess?: (data: TData | null, meta?: PaginatedMeta) => void;
  onError?: (error: ApiError) => void;
};

export function useAction<TInput, TData>(
  action: (input: TInput) => Promise<BaseResponse<TData>>,
  options: ActionOptions<TData> = {}
) {
  // Keep callbacks stable
  const optionsRef = useRef(options);
  useEffect(() => {
    optionsRef.current = options;
  }, [options]);

  const initialState: BaseResponse<TData> = {
    data: null,
    error: null,
  };

  // Wrapper that conforms to useActionState's (prevState, payload) signature
  const wrappedAction = async (
    prevState: BaseResponse<TData>,
    input: TInput
  ): Promise<BaseResponse<TData>> => {
    try {
      const result = await action(input);

      // Check if the API returned an expected error format
      if (result.error) {
        optionsRef.current.onError?.(result.error);
        return result; 
      }

      // Success flow
      optionsRef.current.onSuccess?.(result.data, result.meta);
      return result;

    } catch (err) {
      const fallbackError: ApiError = {
        code: 500,
        message: err instanceof Error ? err.message : "An unexpected error occurred",
      };
      
      optionsRef.current.onError?.(fallbackError);
      return { data: null, error: fallbackError };
    }
  };

  const [state, execute, isPending] = useActionState(wrappedAction, initialState);

  return {
    execute,
    isPending,
    data: state.data,
    error: state.error,
    meta: state.meta,
  };
}